# 📦 Consulta de CEP com Go

Este é um projeto simples em Go que consulta dados de endereço a partir de um CEP usando a API do ViaCEP. Ele também é uma ótima base para quem está aprendendo sobre:

- Criação de servidores HTTP com Go  
- Manipulação de rotas e parâmetros na URL  
- Structs com tags JSON  
- Requisições HTTP externas  
- Conversão entre JSON e structs  
- Tratamento de erros

---

## 🚀 Como executar

### Pré-requisitos:
- [Go instalado](https://go.dev/dl/) (versão 1.16 ou superior)

### Passos:

1. Clone este repositório ou copie o código-fonte.
2. No terminal, vá até a pasta do projeto.
3. Execute:

```bash
go run main.go
```

4. No navegador ou no Postman, acesse:

```
http://localhost:8080/?cep=01001000
```

---

## 🧠 Conceitos importantes aprendidos

### 🌐 Criando um servidor HTTP

```go
http.ListenAndServe(":8080", nil)
```

Esse comando inicia um servidor web escutando na porta 8080.

---

### ⚙️ Definindo a rota principal

```go
http.HandleFunc("/", BuscaCepHandler)
```

A função `BuscaCepHandler` é responsável por lidar com todas as requisições feitas à raiz do site.

---

### 🔎 Recebendo o parâmetro da URL

```go
cepParam := r.URL.Query().Get("cep")
```

Aqui capturamos o valor do parâmetro `cep` informado na URL como uma query string, por exemplo: `/?cep=01001000`.

---

### 🧱 Struct com tags JSON

```go
type ViaCEP struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	// demais campos...
}
```

Essas tags permitem que o Go saiba como mapear os dados do JSON para os campos da struct.

---

### 📤 Fazendo requisição externa para a API ViaCEP

```go
http.Get("https://viacep.com.br/ws/" + cep + "/json/")
```

Isso envia uma requisição GET à API pública do ViaCEP com o CEP informado.

---

### 🧩 Convertendo JSON para struct

```go
json.Unmarshal(body, &c)
```

Usamos `Unmarshal` para transformar o JSON recebido em uma struct do tipo `ViaCEP`.

---

### 📦 Retornando JSON na resposta da API

```go
json.NewEncoder(w).Encode(cep)
```

Isso converte a struct de volta para JSON e envia como resposta para quem fez a requisição.

---

### ⚠️ Tratamento de erros

O código faz verificação de erros após cada operação importante (requisição, leitura e conversão de dados), retornando status HTTP adequados:

- `400 Bad Request`: quando o parâmetro `cep` está ausente.
- `404 Not Found`: quando a rota é diferente de `/`.
- `500 Internal Server Error`: quando ocorre falha na comunicação com a API ou na conversão dos dados.

---

## ✅ Exemplo de resposta JSON

```json
{
  "cep": "01001-000",
  "logradouro": "Praça da Sé",
  "complemento": "lado ímpar",
  "bairro": "Sé",
  "localidade": "São Paulo",
  "uf": "SP",
  "ibge": "3550308",
  "gia": "1004",
  "ddd": "11",
  "siafi": "7107"
}
```

---

## 🛠 Sugestões de melhorias

- Validar se o CEP é válido antes de enviar para a API externa.
- Tratar resposta da API quando o CEP não for encontrado.
- Criar interface gráfica simples para digitação do CEP.
- Adicionar testes automatizados com `httptest`.

---
