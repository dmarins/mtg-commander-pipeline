// Package scryfall baixa e percorre os arquivos bulk data do Scryfall.
//
// O Scryfall publica os dumps em JSONL comprimido com gzip (um objeto JSON por
// linha), o que permite percorrer os ~30k registros de oracle-cards em streaming,
// sem carregar os ~140 MB descomprimidos na memória.
package scryfall

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const bulkIndexURL = "https://api.scryfall.com/bulk-data"

// userAgent identifica o cliente. O Scryfall pede User-Agent e Accept explícitos.
const userAgent = "mtg-commander-pipeline/1.0 (+https://github.com/dmarins/mtg-commander-pipeline)"

// BulkInfo é uma entrada do índice de bulk data.
type BulkInfo struct {
	Type           string `json:"type"`
	Name           string `json:"name"`
	UpdatedAt      string `json:"updated_at"`
	CompressedSize int64  `json:"compressed_size"`
	// A API expõe o dump como JSONL desde 2025; o campo download_uri antigo não existe mais.
	JSONLDownloadURI string `json:"jsonl_download_uri"`
}

func newClient() *http.Client {
	return &http.Client{Timeout: 30 * time.Minute}
}

func get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "*/*")
	resp, err := newClient().Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GET %s: status %s", url, resp.Status)
	}
	return resp, nil
}

// Index devolve o catálogo de bulk data disponível, indexado por Type.
func Index() (map[string]BulkInfo, error) {
	resp, err := get(bulkIndexURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payload struct {
		Data []BulkInfo `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("decodificando índice de bulk data: %w", err)
	}

	out := make(map[string]BulkInfo, len(payload.Data))
	for _, b := range payload.Data {
		out[b.Type] = b
	}
	return out, nil
}

// Stream baixa o dump e chama fn para cada linha JSON, sem materializar o
// arquivo inteiro. As linhas são entregues cruas para que o chamador decodifique
// só os campos que lhe interessam.
func Stream(info BulkInfo, fn func(line []byte) error) error {
	if info.JSONLDownloadURI == "" {
		return fmt.Errorf("bulk %q não tem jsonl_download_uri", info.Type)
	}

	resp, err := get(info.JSONLDownloadURI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("abrindo gzip de %s: %w", info.Type, err)
	}
	defer gz.Close()

	sc := bufio.NewScanner(gz)
	// Cartas com muito texto (e as entradas de tag, que carregam milhares de
	// taggings) passam do buffer padrão de 64 KB.
	sc.Buffer(make([]byte, 0, 1<<20), 64<<20)

	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		if err := fn(line); err != nil {
			return err
		}
	}
	if err := sc.Err(); err != nil && err != io.EOF {
		return fmt.Errorf("lendo %s: %w", info.Type, err)
	}
	return nil
}
