package main

import (
	"log"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/thedivinez/go-libs/gothex"
	"github.com/thedivinez/pocketvault.ai/api"
	"github.com/thedivinez/pocketvault.ai/config"
	"github.com/thedivinez/pocketvault.ai/pages"
)

type Server struct {
	configs *config.Config
	auth    *api.AuthHandler
}

func NewServer() *Server {
	srv := &Server{configs: &config.Config{}, auth: api.NewAuthHandler()}
	if err := srv.configs.ReadFromEnv(); err != nil {
		log.Fatal(err)
	}
	goth.UseProviders(google.New(srv.configs.GOOGLE_KEY, srv.configs.GOOGLE_SECRET, srv.configs.GOOGLE_CALLBACK))
	return srv
}

func main() {
	server := NewServer()
	router := gothex.NewGothexRouter().WithCustomErrorPageContent(
		gothex.ErrorPageContent{
			Code:      404,
			Title:     "Error 404",
			ErrorType: "Resource not found",
			Message:   "The requested resource could not be found but may be available again in the future.",
		},
		gothex.ErrorPageContent{
			Code:      500,
			Title:     "Error 500",
			ErrorType: "Internal Server Error",
			Message:   "An unexpected condition was encountered and prevented the request from succeeding.",
		},
	).WithNoCache()

	gothic.Store = router.CookieStore
	router.GET("", pages.HandleIndex)
	router.GET("/signin", pages.HandleSignIn)
	router.GET("/signup", pages.HandleSignUp)
	router.GET("/deposit", pages.HandleDeposit)
	router.GET("/referrals", pages.HandleReferrals)
	router.GET("/withdraw", pages.HandleWithdrawal)
	router.GET("/two-factor", pages.HandleTwoFactor)
	router.GET("/investments", pages.HandleInvestments)
	router.GET("/ranking", pages.HandleInvestmentRanking)
	router.GET("/transactions", pages.HandleTransactions)
	router.GET("/profile-setting", pages.HandleProfileSetting)
	router.GET("/change-password", pages.HandleChangePassword)
	router.GET("/deposit/history", pages.HandleDepositHistory)
	router.GET("/transfer-balance", pages.HandleTransferBalance)
	router.GET("/invest/schedule", pages.HandleInvestmentSchedule)

	api := router.Group("/api")
	{
		api.Any("/signin", server.auth.HandleSignIn)
		api.POST("/signup", server.auth.HandleSignUp)
		api.GET("/signin/callback", server.auth.HandleProvidersCallback)
		api.POST("/signout", server.auth.HandleSignOut, router.Protected)
	}

	if err := router.Run(); err != nil {
		log.Fatal("failed to start server", err)
	}
}
