package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Struct ViaCEP define o formato esperado da resposta API ViaCep
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// HANDLER QUE RESPONDE AS REQUISIÇÕES FEITAS AO SERVIDOR
// ELE TRATA A ROTA, BUSCA CEP, LIDA COM ERROS E RETORNA A RESPOSTA EM JSON.
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	// Se o caminho da URL for diferente de "/", retorna 404 (página não encontrada)
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// CHAMADA A FUNCAO QUE CONSULTA CEP NA API VIACEP
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest) // Retorna erro 400 (requisição malformada)
		return
	}

	// CHAMA FUNCAO QUE CONSULTA O CEP NA API VIACEP
	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError) // Retorna erro 500 se algo deu errado
		return
	}
	// Define o tipo de resposta como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Serializa automaticamente a struct e escreve no body da resposta
	json.NewEncoder(w).Encode(cep)

}

// Função que consulta a API externa ViaCEP e retorna os dados do endereço.
func BuscaCep(cep string) (*ViaCEP, error) {
	// Faz a requisição GET para a API ViaCEP com o CEP informado
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error // Retorna erro se a requisição falhar
	}
	defer resp.Body.Close() // Garante que o body será fechado após uso

	// Lê todo o conteúdo do body da resposta
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error // Erro ao ler o corpo da resposta
	}

	var c ViaCEP
	// Converte o JSON da resposta para a struct ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error // Erro na conversão JSON → struct
	}

	return &c, nil // Retorna um ponteiro para a struct preenchida
}

// FUNCAO MAIN INICIA O SERVIDOR HTTP LOCAL NA PORTA 8080
func main() {
	http.HandleFunc("/", BuscaCepHandler) // Define que a função handler será chamada quando acessarem "/"
	http.ListenAndServe(":8080", nil)     // Inicia o servidor HTTP na porta 8080
}
