// I BOLETIM_DE_URNA

// NOTAÇÃO: bweb_ <Turno>t_ <UF><Data><Hash> .zip

// Variável								Descrição

// DT_GERACAO							Data de geração do arquivo (data da extração dos dados)
// HH_GERACAO							Hora de geração do arquivo (hora da extração dos dados)
// ANO_ELEICAO							Ano da eleição (referente ao ano eleitoral de pesquisa).
// CD_PLEITO							Código do pleito da eleição.
// DT_PLEITO							Data do pleito da eleição.
// NR_TURNO								Número do turno da eleição.
// CD_ELEICAO							Código da eleição.
// DS_ELEICAO							Descrição da eleição.
// DT_ELEICAO							Data em que ocorreu a eleição.
// SG_UF								Sigla da Unidade da Federação em que ocorreu a eleição.
// CD_MUNICIPIO							Código TSE do município onde ocorreu a eleição.
// NM_MUNICIPIO							Nome do município onde ocorreu a eleição.
// NR_ZONA								Número da Zona Eleitoral em que ocorreu a eleição.
// NR_SECAO								Número da Seção Eleitoral em que ocorreu a eleição.
// NR_LOCAL_VOTACAO						Número do local de votação referente ao boletim de urna.
// CD_CARGO_PERGUNTA					Código do cargo do candidato ou pergunta, no caso de plebiscito.
// DS_CARGO_PERGUNTA					Descrição do cargo do candidato ou pergunta, no caso de plebiscito.
// NR_PARTIDO							Número de identificação do partido.
// SG_PARTIDO							Sigla do partido.
// NM_PARTIDO							Nome do partido.
// QT_APTOS								Quantidade de eleitores aptos a votar naquele município, zona e seção.
// QT_COMPARECIMENTO					Quantidade de eleitores que compareceram às eleições naquele município, zona e seção.
// QT_ABSTENCOES						Quantidade de eleitores que não compareceram às eleições naquele município, zona e seção.
// CD_TIPO_URNA							Código de identificação do tipo de urna.
// DS_TIPO_URNA							Descrição de identificação do tipo de urna.
// CD_TIPO_VOTAVEL						Código do Tipo de votável.
// DS_TIPO_VOTAVEL						Descrição do tipo votavel.
// NR_VOTAVEL							Pode assumir os valores: Número do candidato (quando voto nominal); Número do partido (quando voto em legenda); Número 95 (quando voto em branco); Número 96 (quando voto nulo); Número 97 (quando voto anulado e apurado em separado) e Número 98(quando voto anulado).
// NM_VOTAVEL							Pode assumir os valores: Nome do candidato (quando voto nominal ou voto anulado); Nome do partido (quando voto em legenda); "Voto em branco"(quando voto em branco); "Voto Nulo"(quando voto nulo) e "Voto anulado e apurado em separado"(quando voto anulado e apurado em separado).
// QT_VOTOS								Quantidade de votos recebidos pelo votável.
// NR_URNA_EFETIVADA					Número da urna efetivada.
// CD_CARGA_1_URNA_EFETIVADA			Código da Carga 1 da urna efetivada.
// CD_CARGA_2_URNA_EFETIVADA			Código da Carga 2 da urna efetivada.
// CD_FLASHCARD_URNA_EFETIVADA			Código do Flashcard de urna efetivada.
// DT_CARGA_URNA_EFETIVADA				Data da carga da urna efetivada.
// DS_CARGO_PERGUNTA_SECAO				Informação de cargo, pergunta e seção. Apresentando os valores agrupados de [CD_CARGO_PERGUNTA][CD_SECAO], caso um dos valores seja nulo, deverá retornar "-"no lugar do código.
// DS_AGREGADAS							Lista das seções agregadas.
// DT_ABERTURA							Data/hora de abertura da urna eletrônica para votação.
// DT_ENCERRAMENTO						Data/hora de encerramento da urna eletrônica para votação.
// QT_ELEITORES_BIOMETRIA_NH			Quantitativo de eleitores com biometria, mas que não foram habilitados por meio dela.
// NR_JUNTA_APURADORA					Número da junta apuradora.
// NR_TURMA_APURADORA					Número da turma apuradora.

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func main() {
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}
