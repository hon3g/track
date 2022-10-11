/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"unicode"

	"github.com/spf13/cobra"
)

// lasershipCmd represents the lasership command
var lasershipCmd = &cobra.Command{
	Use:     "lasership <tracking number>",
	Aliases: []string{"ls"},
	Short:   "A brief description of your command",
	Long:    ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if len(args[0]) != 10 || args[0][:2] != "LS" {
			return errors.New("invalid tracking number")
		}
		for _, c := range args[0][2:] {
			if !unicode.IsDigit(c) {
				return errors.New("invalid tracking number")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		printResult(getResult(args[0]), all)
	},
}

func init() {
	rootCmd.AddCommand(lasershipCmd)
	lasershipCmd.Flags().BoolP("all", "a", false, "Get all the events")
}

type ResponseBody struct {
	OrderNumber           string `json:"OrderNumber"`
	ReceivedOn            string `json:"ReceivedOn"`
	UTCReceivedOn         string `json:"UTCReceivedOn"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate"`
	Origin                struct {
		City       string `json:"City"`
		State      string `json:"State"`
		PostalCode string `json:"PostalCode"`
		Country    string `json:"Country"`
	} `json:"Origin"`
	Destination struct {
		City       string `json:"City"`
		State      string `json:"State"`
		PostalCode string `json:"PostalCode"`
		Country    string `json:"Country"`
	} `json:"Destination"`
	Pieces []struct {
		TrackingNumber string  `json:"TrackingNumber"`
		Weight         float64 `json:"Weight"`
		WeightUnit     string  `json:"WeightUnit"`
	} `json:"Pieces"`
	Events []struct {
		DateTime       string `json:"DateTime"`
		UTCDateTime    string `json:"UTCDateTime"`
		City           string `json:"City"`
		State          string `json:"State"`
		PostalCode     string `json:"PostalCode"`
		Country        string `json:"Country"`
		EventType      string `json:"EventType"`
		EventModifier  string `json:"EventModifier"`
		EventLabel     string `json:"EventLabel"`
		EventShortText string `json:"EventShortText"`
		EventLongText  string `json:"EventLongText"`
		Signature      string `json:"Signature"`
		Signature2     string `json:"Signature2"`
		Location       string `json:"Location"`
		Reason         string `json:"Reason"`
	} `json:"Events"`
}

func getResult(trackNum string) ResponseBody {
	base := "https://t.lasership.com/Track/%s/json"
	target := fmt.Sprintf(base, trackNum)
	response, err := http.Get(target)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var res ResponseBody
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalln(err)
	}
	return res
}

func printResult(res ResponseBody, all bool) {
	fmt.Println()
	fmt.Println("Estimated Delivery Date:")
	fmt.Println("  ", res.EstimatedDeliveryDate)

	fmt.Println("Event(s):")
	for _, event := range res.Events {
		fmt.Println("  ", event.DateTime)
		if event.State != "" {
			fmt.Println("  ", event.State, event.City, event.PostalCode)
		}
		fmt.Println("  ", event.EventLongText)
		fmt.Println()
		if !all {
			break
		}
	}
}
