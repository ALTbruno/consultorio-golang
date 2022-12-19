package service

import (
	"fmt"
	"log"
	"regexp"

	"github.com/ALTbruno/consultorio-golang/internal/dto"
	"github.com/ALTbruno/consultorio-golang/internal/model"
	"github.com/ALTbruno/consultorio-golang/internal/repository"
	"github.com/go-playground/validator/v10"
)

func CadastrarConsulta(consulta dto.Consulta) (dto.ConsultaResponse, string) {
	if !DataHoraValida(consulta.DataHora) {
		return dto.ConsultaResponse{}, "Parâmetro dataHora com formato inválido. O formato aceito é: DD-MM-AAAA hh:mm"
	}
	validate := validator.New()
	err := validate.Struct(consulta)
	if err != nil {
		return dto.ConsultaResponse{}, err.Error()
	}
	_, nd := BuscarDentistaPorID(consulta.DentistaID)
	if nd != "" {
		return dto.ConsultaResponse{}, nd
	}
	_, np := BuscarPacientePorID(consulta.PacienteID)
	if np != "" {
		return dto.ConsultaResponse{}, np
	}
	c := repository.CadastrarConsulta(model.Consulta{
		DentistaID: consulta.DentistaID,
		PacienteID: consulta.PacienteID,
		DataHora: consulta.DataHora,
		Descricao: consulta.Descricao,
	})
	consultaResponse := ConverterConsultaModelParaResponse(c)
	return consultaResponse, ""
}

func CadastrarConsultaPorMatriculaRG(consulta dto.ConsultaMatriculaRG) (dto.ConsultaResponse, string) {
	if !DataHoraValida(consulta.DataHora) {
		return dto.ConsultaResponse{}, "Parâmetro dataHora com formato inválido. O formato aceito é: DD-MM-AAAA hh:mm"
	}
	validate := validator.New()
	err := validate.Struct(consulta)
	if err != nil {
		return dto.ConsultaResponse{}, err.Error()
	}
	dentista, nd := BuscarDentistaPorMatricula(consulta.MatriculaDentista)
	if nd != "" {
		return dto.ConsultaResponse{}, nd
	}
	paciente, np := BuscarPacientePorRG(consulta.RGPaciente)
	if np != "" {
		return dto.ConsultaResponse{}, np
	}
	consultaModel := model.Consulta {
		DentistaID: dentista.ID,
		PacienteID: paciente.ID,
		DataHora: consulta.DataHora,
		Descricao: consulta.Descricao,
	}
	c := repository.CadastrarConsulta(consultaModel)
	consultaResponse := ConverterConsultaModelParaResponse(c)
	return consultaResponse, ""
}

func BuscarConsultaPorID(id uint) (dto.ConsultaResponse, string) {
	if !repository.ExisteConsultaPorId(id) {
		return dto.ConsultaResponse{}, fmt.Sprintf("Consulta não encontrada %d", id)
	}
	consulta := repository.BuscarConsultaPorID(id)
	consultaResponse := ConverterConsultaModelParaResponse(consulta)
	return consultaResponse, ""
}

func BuscarConsultasPorRGPaciente(rg string) ([]dto.ConsultaResponse, string) {
	paciente, s := BuscarPacientePorRG(rg)
	if s != "" {
		var consultas []dto.ConsultaResponse
		return consultas, s
	}
	consultas := repository.BuscarConsultasPorIdPaciente(paciente.ID)
	var consultasResponse []dto.ConsultaResponse
	for _, consulta := range consultas {
		consultaResponse := ConverterConsultaModelParaResponse(consulta)
		consultasResponse = append(consultasResponse, consultaResponse)
	}
	return consultasResponse, ""
}

func AtualizarConsultaPorID(consulta dto.Consulta, id uint) (dto.ConsultaResponse, string) {
	if !DataHoraValida(consulta.DataHora) {
		return dto.ConsultaResponse{}, `Parâmetro dataHora com formato inválido. O formato aceito é: "DD-MM-AAAA hh:mm"`
	}
	validate := validator.New()
	err := validate.Struct(consulta)
	if err != nil {
		return dto.ConsultaResponse{}, err.Error()
	}
	_, nd := BuscarDentistaPorID(consulta.DentistaID)
	if nd != "" {
		return dto.ConsultaResponse{}, nd
	}
	_, np := BuscarPacientePorID(consulta.PacienteID)
	if np != "" {
		return dto.ConsultaResponse{}, np
	}
	_, s := BuscarConsultaPorID(id)
	if s != "" {
		return dto.ConsultaResponse{}, s
	}
	colunas := make(map[string]interface{})
	colunas["dentista_id"] = consulta.DentistaID
	colunas["paciente_id"] = consulta.PacienteID
	colunas["data_hora"] = consulta.DataHora
	colunas["descricao"] = consulta.Descricao
	c := repository.AtualizarConsultaPorId(id, colunas)
	consultaResponse := ConverterConsultaModelParaResponse(c)
	return consultaResponse, ""
}

func AtualizarConsultaParcial(consulta dto.Consulta, id uint) (dto.ConsultaResponse, string) {
	_, s := BuscarConsultaPorID(id)
	if s != "" {
		return dto.ConsultaResponse{}, s
	}
	colunas := make(map[string]interface{})
	if consulta.DentistaID != 0 {
		_, s := BuscarDentistaPorID(consulta.DentistaID)
		if s != "" {
			return dto.ConsultaResponse{}, s
		}
		colunas["dentista_id"] = consulta.DentistaID
	}
	if consulta.PacienteID != 0 {
		_, s := BuscarPacientePorID(consulta.PacienteID)
		if s != "" {
			return dto.ConsultaResponse{}, s
		}
		colunas["paciente_id"] = consulta.PacienteID
	}
	if consulta.DataHora != "" {
		if !DataHoraValida(consulta.DataHora) {
			return dto.ConsultaResponse{}, "Parâmetro dataHora com formato inválido. O formato aceito é: DD-MM-AAAA hh:mm"
		}
		colunas["data_hora"] = consulta.DataHora
	}
	if consulta.Descricao != "" {
		colunas["descricao"] = consulta.Descricao
	}
	
	c := repository.AtualizarConsultaPorId(id, colunas)
	consultaResponse := ConverterConsultaModelParaResponse(c)
	return consultaResponse, ""
}

func DeletarConsultaPorID(id uint) (int, string) {
	_, res := BuscarConsultaPorID(id)
	if res != "" {
		return 400, res
	}
	result := repository.DeletarConsultaPorId(id)
	if result {
		return 200, fmt.Sprintf("Consulta %d deletada com sucesso", id)
	} else {
		return 500, fmt.Sprintf("Erro ao deletar a consulta %d", id)
	}
}

func DataHoraValida(dataHora string) bool {
	pattern := `^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-(\d{4}) (00|[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$`
	matched, err := regexp.MatchString(pattern, dataHora)
	if err != nil {
		log.Fatal("Erro ao validar Data e Hora com o padrão")
	}
	return matched
}

func ConverterConsultaModelParaResponse(consulta model.Consulta) dto.ConsultaResponse {
	return dto.ConsultaResponse{
		ID: consulta.ID,
		DentistaID: consulta.DentistaID,
		PacienteID: consulta.PacienteID,
		DataHora: consulta.DataHora,
		Descricao: consulta.Descricao,
	}
}
