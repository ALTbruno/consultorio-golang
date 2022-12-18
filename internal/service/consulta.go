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

func CadastrarConsulta(consulta model.Consulta) (model.Consulta, string) {
	if !DataHoraValida(consulta.DataHora) {
		return model.Consulta{}, "Parâmetro dataHora com formato inválido. O formato aceito é: DD-MM-AAAA hh:mm"
	}
	validate := validator.New()
	err := validate.Struct(consulta)
	if err != nil {
		return model.Consulta{}, err.Error()
	}
	_, nd := BuscarDentistaPorID(int(consulta.DentistaID))
	if nd != "" {
		return model.Consulta{}, nd
	}
	_, np := BuscarPacientePorID(int(consulta.PacienteID))
	if np != "" {
		return model.Consulta{}, np
	}
	c := repository.CadastrarConsulta(consulta)
	return c, ""
}

func CadastrarConsultaPorMatriculaRG(consulta dto.ConsultaMatriculaRG) (model.Consulta, string) {
	if !DataHoraValida(consulta.DataHora) {
		return model.Consulta{}, "Parâmetro dataHora com formato inválido. O formato aceito é: DD-MM-AAAA hh:mm"
	}
	validate := validator.New()
	err := validate.Struct(consulta)
	if err != nil {
		return model.Consulta{}, err.Error()
	}
	dentista, nd := BuscarDentistaPorMatricula(consulta.MatriculaDentista)
	if nd != "" {
		return model.Consulta{}, nd
	}
	paciente, np := BuscarPacientePorRG(consulta.RGPaciente)
	if np != "" {
		return model.Consulta{}, np
	}
	consultaModel := model.Consulta {
		DentistaID: dentista.ID,
		PacienteID: paciente.ID,
		DataHora: consulta.DataHora,
		Descricao: consulta.Descricao,
	}
	c := repository.CadastrarConsulta(consultaModel)
	return c, ""
}

func BuscarConsultaPorID(id int) (model.Consulta, string) {
	if !repository.ExisteConsultaPorId(id) {
		return model.Consulta{}, fmt.Sprintf("Consulta não encontrada %d", id)
	}
	consulta := repository.BuscarConsultaPorID(id)
	return consulta, ""
}

func AtualizarConsultaPorID(consulta model.Consulta, id int) (model.Consulta, string) {
	if !DataHoraValida(consulta.DataHora) {
		return model.Consulta{}, `Parâmetro dataHora com formato inválido. O formato aceito é: "DD-MM-AAAA hh:mm"`
	}
	validate := validator.New()
	err := validate.Struct(consulta)
	if err != nil {
		return model.Consulta{}, err.Error()
	}
	_, nd := BuscarDentistaPorID(int(consulta.DentistaID))
	if nd != "" {
		return model.Consulta{}, nd
	}
	_, np := BuscarPacientePorID(int(consulta.PacienteID))
	if np != "" {
		return model.Consulta{}, np
	}
	consultaSalva, s := BuscarConsultaPorID(id)
	if s != "" {
		return model.Consulta{}, s
	}
	consultaSalva.DentistaID = consulta.DentistaID
	consultaSalva.PacienteID = consulta.PacienteID
	consultaSalva.DataHora = consulta.DataHora
	consultaSalva.Descricao = consulta.Descricao
	c := repository.AtualizarConsulta(consultaSalva)
	return c, ""
}

func AtualizarConsultaParcial(consulta model.Consulta, id int) (model.Consulta, string) {
	consultaSalva, s := BuscarConsultaPorID(id)
	if s != "" {
		return model.Consulta{}, s
	}
	colunas := make(map[string]interface{})
	if consulta.DentistaID != 0 {
		_, s := BuscarDentistaPorID(int(consulta.DentistaID))
		if s != "" {
			return model.Consulta{}, s
		}
		colunas["dentista_id"] = consulta.DentistaID
	}
	if consulta.PacienteID != 0 {
		_, s := BuscarPacientePorID(int(consulta.PacienteID))
		if s != "" {
			return model.Consulta{}, s
		}
		colunas["paciente_id"] = consulta.PacienteID
	}
	if consulta.DataHora != "" {
		if !DataHoraValida(consulta.DataHora) {
			return model.Consulta{}, "Parâmetro dataHora com formato inválido. O formato aceito é: DD-MM-AAAA hh:mm"
		}
		colunas["data_hora"] = consulta.DataHora
	}
	if consulta.Descricao != "" {
		colunas["descricao"] = consulta.Descricao
	}
	
	c := repository.AtualizarConsultaParcial(consultaSalva, colunas)
	return c, ""
}

func DeletarConsultaPorID(id int) (int, string) {
	consulta, res := BuscarConsultaPorID(id)
	if res != "" {
		return 400, res
	}
	result := repository.DeletarConsulta(consulta)
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
