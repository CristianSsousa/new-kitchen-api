package controllers

import (
	"net/http"

	"cha-casa-nova-backend/database"
	"cha-casa-nova-backend/models"

	"github.com/gin-gonic/gin"
)

// GetStats retorna estatísticas gerais do sistema
func GetStats(c *gin.Context) {
	var stats struct {
		TotalItens           int64   `json:"total_itens"`
		ItensResgatados      int64   `json:"itens_resgatados"`
		TotalMensagens       int64   `json:"total_mensagens"`
		TotalConvidados      int     `json:"total_convidados"`
		PorcentagemConcluida float64 `json:"porcentagem_concluida"`
	}

	// Total de itens
	if err := database.DB.Model(&models.Item{}).Count(&stats.TotalItens).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de itens"})
		return
	}

	// Itens resgatados
	if err := database.DB.Model(&models.Item{}).Where("resgatado = ?", true).Count(&stats.ItensResgatados).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular itens resgatados"})
		return
	}

	// Total de mensagens
	if err := database.DB.Model(&models.Mensagem{}).Where("aprovada = ?", true).Count(&stats.TotalMensagens).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de mensagens"})
		return
	}

	// Total de convidados
	var totalAdultos, totalCriancas int64
	if err := database.DB.Model(&models.Confirmacao{}).Select("COALESCE(SUM(quantidade_adultos), 0)").Row().Scan(&totalAdultos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de adultos"})
		return
	}
	if err := database.DB.Model(&models.Confirmacao{}).Select("COALESCE(SUM(quantidade_criancas), 0)").Row().Scan(&totalCriancas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de crianças"})
		return
	}
	stats.TotalConvidados = int(totalAdultos + totalCriancas)

	// Porcentagem concluída
	if stats.TotalItens > 0 {
		stats.PorcentagemConcluida = float64(stats.ItensResgatados) / float64(stats.TotalItens) * 100
	}

	c.JSON(http.StatusOK, stats)
}

// GetEstatisticasDetalhadas retorna estatísticas mais detalhadas do sistema
func GetEstatisticasDetalhadas(c *gin.Context) {
	// Estatísticas básicas
	var stats struct {
		TotalItens           int64   `json:"total_itens"`
		ItensResgatados      int64   `json:"itens_resgatados"`
		TotalMensagens       int64   `json:"total_mensagens"`
		TotalConvidados      int64   `json:"total_convidados"`
		TotalConfirmacoes    int64   `json:"total_confirmacoes"`
		ValorTotalItens      float64 `json:"valor_total_itens"`
		PorcentagemConcluida float64 `json:"porcentagem_concluida"`
		MensagensPendentes   int64   `json:"mensagens_pendentes"`
		CategoriaStats       []struct {
			Categoria   string  `json:"categoria"`
			Total       int64   `json:"total"`
			Resgatados  int64   `json:"resgatados"`
			Porcentagem float64 `json:"porcentagem"`
		} `json:"categorias"`
	}

	// Total de itens
	if err := database.DB.Model(&models.Item{}).Count(&stats.TotalItens).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de itens"})
		return
	}

	// Itens resgatados
	if err := database.DB.Model(&models.Item{}).Where("resgatado = ?", true).Count(&stats.ItensResgatados).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular itens resgatados"})
		return
	}

	// Total de mensagens
	if err := database.DB.Model(&models.Mensagem{}).Where("aprovada = ?", true).Count(&stats.TotalMensagens).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de mensagens"})
		return
	}

	// Mensagens pendentes
	if err := database.DB.Model(&models.Mensagem{}).Where("aprovada = ?", false).Count(&stats.MensagensPendentes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular mensagens pendentes"})
		return
	}

	// Total de confirmações
	if err := database.DB.Model(&models.Confirmacao{}).Count(&stats.TotalConfirmacoes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de confirmações"})
		return
	}

	// Total de convidados
	var totalAdultos, totalCriancas int64
	if err := database.DB.Model(&models.Confirmacao{}).Select("COALESCE(SUM(quantidade_adultos), 0)").Row().Scan(&totalAdultos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de adultos"})
		return
	}
	if err := database.DB.Model(&models.Confirmacao{}).Select("COALESCE(SUM(quantidade_criancas), 0)").Row().Scan(&totalCriancas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular total de crianças"})
		return
	}
	stats.TotalConvidados = totalAdultos + totalCriancas

	// Valor total dos itens
	if err := database.DB.Model(&models.Item{}).Select("COALESCE(SUM(preco), 0)").Row().Scan(&stats.ValorTotalItens); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular valor total dos itens"})
		return
	}

	// Porcentagem concluída
	if stats.TotalItens > 0 {
		stats.PorcentagemConcluida = float64(stats.ItensResgatados) / float64(stats.TotalItens) * 100
	}

	// Estatísticas por categoria
	rows, err := database.DB.Raw(`
		SELECT 
			categoria,
			COUNT(*) as total,
			SUM(CASE WHEN resgatado = true THEN 1 ELSE 0 END) as resgatados
		FROM items 
		WHERE deleted_at IS NULL
		GROUP BY categoria
	`).Rows()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular estatísticas por categoria"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var catStat struct {
			Categoria   string  `json:"categoria"`
			Total       int64   `json:"total"`
			Resgatados  int64   `json:"resgatados"`
			Porcentagem float64 `json:"porcentagem"`
		}
		if err := rows.Scan(&catStat.Categoria, &catStat.Total, &catStat.Resgatados); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler estatísticas por categoria"})
			return
		}
		if catStat.Total > 0 {
			catStat.Porcentagem = float64(catStat.Resgatados) / float64(catStat.Total) * 100
		}
		stats.CategoriaStats = append(stats.CategoriaStats, catStat)
	}

	c.JSON(http.StatusOK, stats)
}
