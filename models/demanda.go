package models

// Demanda objeto de retorno da demanda do acelerato
type Demanda struct {
	ID                     int    `json:"id"`
	TicketKey              int    `json:"ticketKey"`
	Arquivado              bool   `json:"arquivado"`
	Lixeira                bool   `json:"lixeira"`
	Suspenso               bool   `json:"suspenso"`
	Mesclado               bool   `json:"mesclado"`
	TempoDeVidaEmDias      int    `json:"tempoDeVidaEmDias,omitempty"`
	TempoCiclico           int    `json:"tempoCiclico,omitempty"`
	QuantidadeDeVinculados int    `json:"quantidadeDeVinculados"`
	Titulo                 string `json:"titulo"`
	Descricao              string `json:"descricao"`
	DataDaUltimaAlteracao  struct {
		Timestamp int64  `json:"timestamp"`
		Formatted string `json:"formatted"`
	} `json:"dataDaUltimaAlteracao"`
	DataDeCriacao struct {
		Timestamp int64  `json:"timestamp"`
		Formatted string `json:"formatted"`
	} `json:"dataDeCriacao"`
	Categoria struct {
		CategoriaKey int    `json:"categoriaKey"`
		Descricao    string `json:"descricao"`
		Cor          string `json:"cor"`
	} `json:"categoria"`
	TipoDeTicket struct {
		TipoDeTicketKey int    `json:"tipoDeTicketKey"`
		Descricao       string `json:"descricao"`
		Cor             string `json:"cor"`
	} `json:"tipoDeTicket,omitempty"`
	KanbanStatus struct {
		KanbanStatusKey int    `json:"kanbanStatusKey"`
		Descricao       string `json:"descricao"`
		Inicio          bool   `json:"inicio"`
		Fim             bool   `json:"fim"`
		Fila            bool   `json:"fila"`
		Ordem           int    `json:"ordem"`
	} `json:"kanbanStatus"`
	Reporter struct {
		UsuarioKey     int    `json:"usuarioKey"`
		AvatarPath     string `json:"avatarPath"`
		Nome           string `json:"nome"`
		EmailPrincipal string `json:"emailPrincipal"`
		UltimoAcessoEm struct {
			Timestamp int64  `json:"timestamp"`
			Formatted string `json:"formatted"`
		} `json:"ultimoAcessoEm"`
		Ativo                bool          `json:"ativo"`
		ApenasEmail          bool          `json:"apenasEmail"`
		Emails               []string      `json:"emails"`
		Telefones            []interface{} `json:"telefones"`
		Tags                 []interface{} `json:"tags"`
		OrganizacaoPrincipal struct {
			OrganizacaoKey   int    `json:"organizacaoKey"`
			Nome             string `json:"nome"`
			NomeParaExibicao string `json:"nomeParaExibicao"`
			Observacoes      string `json:"observacoes"`
			Codigo           string `json:"codigo"`
			Cnpj             string `json:"cnpj"`
			Telefone         string `json:"telefone"`
			Ativo            bool   `json:"ativo"`
			Equipe           struct {
				EquipeKey    int    `json:"equipeKey"`
				Nome         string `json:"nome"`
				TipoDeEquipe string `json:"tipoDeEquipe"`
			} `json:"equipe"`
			ValorHora    int           `json:"valorHora"`
			HorasSuporte int           `json:"horasSuporte"`
			Visualizavel bool          `json:"visualizavel"`
			OrgDomains   []interface{} `json:"orgDomains"`
		} `json:"organizacaoPrincipal"`
		Organizacoes []struct {
			OrganizacaoKey   int    `json:"organizacaoKey"`
			Nome             string `json:"nome"`
			NomeParaExibicao string `json:"nomeParaExibicao"`
			Observacoes      string `json:"observacoes"`
			Codigo           string `json:"codigo"`
			Cnpj             string `json:"cnpj"`
			Telefone         string `json:"telefone"`
			Ativo            bool   `json:"ativo"`
			Equipe           struct {
				EquipeKey    int    `json:"equipeKey"`
				Nome         string `json:"nome"`
				TipoDeEquipe string `json:"tipoDeEquipe"`
			} `json:"equipe"`
			ValorHora    int           `json:"valorHora"`
			HorasSuporte int           `json:"horasSuporte"`
			Visualizavel bool          `json:"visualizavel"`
			OrgDomains   []interface{} `json:"orgDomains"`
		} `json:"organizacoes"`
		RedesSociais []interface{} `json:"redesSociais"`
	} `json:"reporter"`
	TipoDePrioridade struct {
		TipoDePrioridadeKey int    `json:"tipoDePrioridadeKey"`
		Descricao           string `json:"descricao"`
		Cor                 string `json:"cor"`
		Peso                int    `json:"peso"`
	} `json:"tipoDePrioridade,omitempty"`
	EspecieDeTicket struct {
		EspecieDeTicketKey int    `json:"especieDeTicketKey"`
		Descricao          string `json:"descricao"`
		Slug               string `json:"slug"`
	} `json:"especieDeTicket"`
	Release struct {
		ReleaseKey int    `json:"releaseKey"`
		Nome       string `json:"nome"`
		Pai        struct {
			ReleaseKey int    `json:"releaseKey"`
			Nome       string `json:"nome"`
			Pai        struct {
				ReleaseKey  int    `json:"releaseKey"`
				Nome        string `json:"nome"`
				DataInicial struct {
					Timestamp int64  `json:"timestamp"`
					Formatted string `json:"formatted"`
				} `json:"dataInicial"`
				DataFinal struct {
					Timestamp int64  `json:"timestamp"`
					Formatted string `json:"formatted"`
				} `json:"dataFinal"`
			} `json:"pai"`
			DataInicial struct {
				Timestamp int64  `json:"timestamp"`
				Formatted string `json:"formatted"`
			} `json:"dataInicial"`
			DataFinal struct {
				Timestamp int64  `json:"timestamp"`
				Formatted string `json:"formatted"`
			} `json:"dataFinal"`
		} `json:"pai"`
		DataInicial struct {
			Timestamp int64  `json:"timestamp"`
			Formatted string `json:"formatted"`
		} `json:"dataInicial"`
		DataFinal struct {
			Timestamp int64  `json:"timestamp"`
			Formatted string `json:"formatted"`
		} `json:"dataFinal"`
	} `json:"release,omitempty"`
	Listas []struct {
		ListaDeTicketsKey int    `json:"listaDeTicketsKey"`
		Nome              string `json:"nome"`
	} `json:"listas,omitempty"`
	Esforco int `json:"esforco,omitempty"`
	Projeto struct {
		ProjetoKey int    `json:"projetoKey"`
		Nome       string `json:"nome"`
		Ativo      bool   `json:"ativo"`
	} `json:"projeto"`
	Responsavel struct {
		UsuarioKey     int    `json:"usuarioKey"`
		AvatarPath     string `json:"avatarPath"`
		Nome           string `json:"nome"`
		EmailPrincipal string `json:"emailPrincipal"`
		UltimoAcessoEm struct {
			Timestamp int64  `json:"timestamp"`
			Formatted string `json:"formatted"`
		} `json:"ultimoAcessoEm"`
		Ativo                bool          `json:"ativo"`
		ApenasEmail          bool          `json:"apenasEmail"`
		Emails               []string      `json:"emails"`
		Telefones            []interface{} `json:"telefones"`
		Tags                 []interface{} `json:"tags"`
		OrganizacaoPrincipal struct {
			OrganizacaoKey   int    `json:"organizacaoKey"`
			Nome             string `json:"nome"`
			NomeParaExibicao string `json:"nomeParaExibicao"`
			Observacoes      string `json:"observacoes"`
			Codigo           string `json:"codigo"`
			Cnpj             string `json:"cnpj"`
			Telefone         string `json:"telefone"`
			Ativo            bool   `json:"ativo"`
			Equipe           struct {
				EquipeKey    int    `json:"equipeKey"`
				Nome         string `json:"nome"`
				TipoDeEquipe string `json:"tipoDeEquipe"`
			} `json:"equipe"`
			ValorHora    int           `json:"valorHora"`
			HorasSuporte int           `json:"horasSuporte"`
			Visualizavel bool          `json:"visualizavel"`
			OrgDomains   []interface{} `json:"orgDomains"`
		} `json:"organizacaoPrincipal"`
		Organizacoes []struct {
			OrganizacaoKey   int    `json:"organizacaoKey"`
			Nome             string `json:"nome"`
			NomeParaExibicao string `json:"nomeParaExibicao"`
			Observacoes      string `json:"observacoes"`
			Codigo           string `json:"codigo"`
			Cnpj             string `json:"cnpj"`
			Telefone         string `json:"telefone"`
			Ativo            bool   `json:"ativo"`
			Equipe           struct {
				EquipeKey    int    `json:"equipeKey"`
				Nome         string `json:"nome"`
				TipoDeEquipe string `json:"tipoDeEquipe"`
			} `json:"equipe"`
			ValorHora    int           `json:"valorHora"`
			HorasSuporte int           `json:"horasSuporte"`
			Visualizavel bool          `json:"visualizavel"`
			OrgDomains   []interface{} `json:"orgDomains"`
		} `json:"organizacoes"`
		RedesSociais []interface{} `json:"redesSociais"`
	} `json:"responsavel,omitempty"`
	URL      string `json:"url"`
	Atuacoes []struct {
		AtuacaoKey           int  `json:"atuacaoKey"`
		Ativo                bool `json:"ativo"`
		DataDaUltimaAtivacao struct {
			Timestamp int64  `json:"timestamp"`
			Formatted string `json:"formatted"`
		} `json:"dataDaUltimaAtivacao"`
		Usuario struct {
			UsuarioKey     int    `json:"usuarioKey"`
			AvatarPath     string `json:"avatarPath"`
			Nome           string `json:"nome"`
			EmailPrincipal string `json:"emailPrincipal"`
			UltimoAcessoEm struct {
				Timestamp int64  `json:"timestamp"`
				Formatted string `json:"formatted"`
			} `json:"ultimoAcessoEm"`
			Ativo                bool          `json:"ativo"`
			ApenasEmail          bool          `json:"apenasEmail"`
			Emails               []string      `json:"emails"`
			Telefones            []interface{} `json:"telefones"`
			Tags                 []interface{} `json:"tags"`
			OrganizacaoPrincipal struct {
				OrganizacaoKey   int    `json:"organizacaoKey"`
				Nome             string `json:"nome"`
				NomeParaExibicao string `json:"nomeParaExibicao"`
				Observacoes      string `json:"observacoes"`
				Codigo           string `json:"codigo"`
				Cnpj             string `json:"cnpj"`
				Telefone         string `json:"telefone"`
				Ativo            bool   `json:"ativo"`
				Equipe           struct {
					EquipeKey    int    `json:"equipeKey"`
					Nome         string `json:"nome"`
					TipoDeEquipe string `json:"tipoDeEquipe"`
				} `json:"equipe"`
				ValorHora    int           `json:"valorHora"`
				HorasSuporte int           `json:"horasSuporte"`
				Visualizavel bool          `json:"visualizavel"`
				OrgDomains   []interface{} `json:"orgDomains"`
			} `json:"organizacaoPrincipal"`
			Organizacoes []struct {
				OrganizacaoKey   int    `json:"organizacaoKey"`
				Nome             string `json:"nome"`
				NomeParaExibicao string `json:"nomeParaExibicao"`
				Observacoes      string `json:"observacoes"`
				Codigo           string `json:"codigo"`
				Cnpj             string `json:"cnpj"`
				Telefone         string `json:"telefone"`
				Ativo            bool   `json:"ativo"`
				Equipe           struct {
					EquipeKey    int    `json:"equipeKey"`
					Nome         string `json:"nome"`
					TipoDeEquipe string `json:"tipoDeEquipe"`
				} `json:"equipe"`
				ValorHora    int           `json:"valorHora"`
				HorasSuporte int           `json:"horasSuporte"`
				Visualizavel bool          `json:"visualizavel"`
				OrgDomains   []interface{} `json:"orgDomains"`
			} `json:"organizacoes"`
			RedesSociais []interface{} `json:"redesSociais"`
		} `json:"usuario"`
	} `json:"atuacoes,omitempty"`
	Tags []struct {
		TicketTagKey int    `json:"ticketTagKey"`
		Tag          string `json:"tag"`
	} `json:"tags,omitempty"`
	DataLimite struct {
		Timestamp int64  `json:"timestamp"`
		Formatted string `json:"formatted"`
		Cor       string `json:"cor"`
		Descricao string `json:"descricao"`
	} `json:"dataLimite,omitempty"`
}
