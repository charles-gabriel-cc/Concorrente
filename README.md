<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&color=0:0c4a6e,50:0284c7,100:00add8&height=170&section=header&text=Programa%C3%A7%C3%A3o%20Concorrente&fontSize=42&fontColor=ffffff&animation=fadeIn&fontAlignY=40&desc=Problemas%20Cl%C3%A1ssicos%20de%20Sincroniza%C3%A7%C3%A3o%20em%20Go%20(Goroutines%20%2B%20Mutex)&descAlignY=65&descSize=15" width="100%" alt="Banner" />
</div>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Goroutines-sync.Mutex-0284C7?style=for-the-badge" alt="Goroutines" />
</p>

---

## 📖 Sobre o Projeto

Coleção de exercícios de **Programação Concorrente** em **Go**, explorando sincronização entre *goroutines*, exclusão mútua com `sync.Mutex`, coordenação com `sync.WaitGroup` e cedência cooperativa com `runtime.Gosched()`.

---

## 🧩 Exercícios Implementados

| Caminho | Problema | Mecanismos Utilizados |
| :--- | :--- | :--- |
| `Exercício 02/Questão 01/main/main.go` | **Produtor-Consumidor** com `MutexEventBuffer` | `sync.Mutex`, `sync.WaitGroup`, `runtime.Gosched()` |
| `Exercício 02/Questão 02/main.go` | **Barbeiro Dorminhoco** (`cliente` / `barbeiro` + buffer limitado) | `sync.Mutex`, `sync.WaitGroup`, `runtime.Gosched()` |

---

## 🚀 Como Executar

Com o [Go](https://go.dev/dl/) instalado:

```bash
# Clone o repositório
git clone https://github.com/charles-gabriel-cc/Concorrente.git
cd Concorrente

# Executar Questão 01 (Produtor-Consumidor)
go run "Exercício 02/Questão 01/main/main.go"

# Executar Questão 02 (Barbeiro Dorminhoco)
go run "Exercício 02/Questão 02/main.go"
```

---

<div align="center">
  <sub>Desenvolvido por <a href="https://github.com/charles-gabriel-cc">Charles Gabriel</a></sub>
</div>
