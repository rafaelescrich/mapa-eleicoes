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
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Election struct holds all information about a election process
type Election struct {
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
	dtGeracao                string `json:"dt_geracao"`
	hhGeracao                string `json:"hh_geracao"`
	anoEleicao               string `json:"ano_eleicao"`
	cdPleito                 string `json:"cd_pleito"`
	dtPleito                 string `json:"dt_pleito"`
	nrTurno                  string `json:"nr_turno"`
	cdEleicao                string `json:"cd_eleicao"`
	dsEleicao                string `json:"ds_eleicao"`
	dtEleicao                string `json:"dt_eleicao"`
	sgUf                     string `json:"sg_uf"`
	cdMunicipio              string `json:"cd_municipio"`
	nmMunicipio              string `json:"nm_municipio"`
	nrZona                   string `json:"nr_zona"`
	nrSecao                  string `json:"nr_secao"`
	nrLocalVotacao           string `json:"nr_local_votacao"`
	cdCargoPergunta          string `json:"cd_cargo_pergunta"`
	dsCargoPergunta          string `json:"ds_cargo_pergunta"`
	nrPartido                string `json:"nr_partido"`
	sgPartido                string `json:"sg_partido"`
	nmPartido                string `json:"nm_partido"`
	qtAptos                  string `json:"qt_aptos"`
	qtComparecimento         string `json:"qt_comparecimento"`
	qtAbstencoes             string `json:"qt_abstencoes"`
	cdTipoUrna               string `json:"cd_tipo_urna"`
	dsTipoUrna               string `json:"ds_tipo_urna"`
	cdTipoVotavel            string `json:"cd_tipo_votavel"`
	dsTipoVotavel            string `json:"ds_tipo_votavel"`
	nrVotavel                string `json:"nr_votavel"`
	nmVotavel                string `json:"nm_votavel"`
	qtVotos                  string `json:"qt_votos"`
	nrUrnaEfetivada          string `json:"nr_urna_efetivada"`
	cdCarga1UrnaEfetivada    string `json:"cd_carga_1_urna_efetivada"`
	cdCarga2UrnaEfetivada    string `json:"cd_carga_2_urna_efetivada"`
	cdFlashcardUrnaEfetivada string `json:"cd_flashcard_urna_efetivada"`
	dtCargaUrnaEfetivada     string `json:"dt_carga_urna_efetivada"`
	dsCargoPerguntaSecao     string `json:"ds_cargo_pergunta_secao"`
	dsAgregadas              string `json:"ds_agregadas"`
	dtAbertura               string `json:"dt_abertura"`
	dtEncerramento           string `json:"dt_encerramento"`
	qtEleitoresBiometriaNh   string `json:"qt_eleitores_biometria_nh"`
	nrJuntaApuradora         string `json:"nr_junta_apuradora"`
	nrTurmaApuradora         string `json:"nr_turma_apuradora"`
}

//type Person struct {
//	Firstname string   `json:"firstname"`
//	Lastname  string   `json:"lastname"`
//	Address   *Address `json:"address,omitempty"`
//}

//type State struct {
//	City  string `json:"city"`
//	State string `json:"state"`
//}

func main() {

	csvFile, _ := os.Open("test.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	//reader.LazyQuotes = true
	reader.TrailingComma = true
	var electionResults []Election
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		electionResults = append(electionResults, Election{
			dtGeracao:                line[0],
			hhGeracao:                line[2],
			anoEleicao:               line[3],
			cdPleito:                 line[4],
			dtPleito:                 line[5],
			nrTurno:                  line[6],
			cdEleicao:                line[7],
			dsEleicao:                line[8],
			dtEleicao:                line[9],
			sgUf:                     line[10],
			cdMunicipio:              line[11],
			nmMunicipio:              line[12],
			nrZona:                   line[13],
			nrSecao:                  line[14],
			nrLocalVotacao:           line[15],
			cdCargoPergunta:          line[16],
			dsCargoPergunta:          line[17],
			nrPartido:                line[18],
			sgPartido:                line[19],
			nmPartido:                line[20],
			qtAptos:                  line[21],
			qtComparecimento:         line[22],
			qtAbstencoes:             line[23],
			cdTipoUrna:               line[24],
			dsTipoUrna:               line[25],
			cdTipoVotavel:            line[26],
			dsTipoVotavel:            line[27],
			nrVotavel:                line[28],
			nmVotavel:                line[29],
			qtVotos:                  line[30],
			nrUrnaEfetivada:          line[31],
			cdCarga1UrnaEfetivada:    line[32],
			cdCarga2UrnaEfetivada:    line[33],
			cdFlashcardUrnaEfetivada: line[34],
			dtCargaUrnaEfetivada:     line[35],
			dsCargoPerguntaSecao:     line[36],
			dsAgregadas:              line[37],
			dtAbertura:               line[38],
			dtEncerramento:           line[39],
			qtEleitoresBiometriaNh:   line[40],
			nrJuntaApuradora:         line[41],
			nrTurmaApuradora:         line[42],
		})
	}
	electionRJSON, _ := json.Marshal(electionResults)
	fmt.Println(string(electionRJSON))
}
