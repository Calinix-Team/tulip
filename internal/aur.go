package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/Jguer/aur"
)

const (
	boldCode  = "\x1b[1m"
	resetCode = "\x1b[0m"
)

const UseColor = true // nolintcolor

const (
	searchMode = "search"
	infoMode   = "info"
)

func getSearchBy(value string) aur.By {
	switch value {
	case "name":
		return aur.Name
	case "maintainer":
		return aur.Maintainer
	case "depends":
		return aur.Depends
	case "makedepends":
		return aur.MakeDepends
	case "optdepends":
		return aur.OptDepends
	case "checkdepends":
		return aur.CheckDepends
	default:
		return aur.NameDesc
	}
}

func versionRequestEditor(ctx context.Context, req *http.Request) error {
	req.Header.Add("User-Agent", "aur-cli/v1")

	return nil
}

func GetResults(by string, aurURL string, verbose bool, jsonDisplay bool, pkgs []string) {

	mode := searchMode

	aurClient, err := aur.NewClient(aur.WithBaseURL(aurURL),
		aur.WithRequestEditorFn(versionRequestEditor))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	results, err := getResults(aurClient, by, mode, pkgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	if jsonDisplay {
		output, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

			os.Exit(1)
		}

		fmt.Println(string(output))
	} else {
		for i := range results {
			switch mode {
			case infoMode:
				printInfo(&results[i], os.Stdout, aurURL, verbose)
			case searchMode:
				printSearch(&results[i], os.Stdout)
			default:
				os.Exit(1)
			}
		}
	}
}

func getResults(aurClient *aur.Client, by, mode string, pkgs []string) (results []aur.Pkg, err error) {
	switch mode {
	case searchMode:
		results, err = aurClient.Search(context.Background(), strings.Join(pkgs, " "), getSearchBy(by))
	case infoMode:
		results, err = aurClient.Info(context.Background(), pkgs)
	default:
		os.Exit(1)
	}

	if err != nil {
		err = fmt.Errorf("rpc request failed: %w", err)
	}

	return results, err
}

func stylize(startCode, in string) string {
	if UseColor {
		return startCode + in + resetCode
	}
	return in
}

func Bold(in string) string {
	return stylize(boldCode, in)
}

func printSearch(a *aur.Pkg, w io.Writer) {
	fmt.Fprintf(w, "- %s %s (%d %.2f)\n\t%s\n",
		Bold(a.Name), a.Version, a.NumVotes, a.Popularity, a.Description)
}

// PrintInfo prints package info like pacman -Si.
func printInfo(pkg *aur.Pkg, writer io.Writer, aurURL string, verbose bool) {
	printInfoValue(writer, "Name", pkg.Name)
	printInfoValue(writer, "Version", pkg.Version)
	printInfoValue(writer, "Description", pkg.Description)

	if verbose {
		printInfoValue(writer, "Keywords", pkg.Keywords...)
		printInfoValue(writer, "URL", pkg.URL)
		printInfoValue(writer, "AUR URL", strings.TrimRight(aurURL, "/")+"/packages/"+pkg.Name)

		printInfoValue(writer, "Groups", pkg.Groups...)
		printInfoValue(writer, "Licenses", pkg.License...)
		printInfoValue(writer, "Provides", pkg.Provides...)
		printInfoValue(writer, "Depends On", pkg.Depends...)
		printInfoValue(writer, "Make Deps", pkg.MakeDepends...)
		printInfoValue(writer, "Check Deps", pkg.CheckDepends...)
		printInfoValue(writer, "Optional Deps", pkg.OptDepends...)
		printInfoValue(writer, "Conflicts With", pkg.Conflicts...)

		printInfoValue(writer, "Maintainer", pkg.Maintainer)
		printInfoValue(writer, "Votes", fmt.Sprintf("%d", pkg.NumVotes))
		printInfoValue(writer, "Popularity", fmt.Sprintf("%f", pkg.Popularity))
		printInfoValue(writer, "First Submitted", formatTimeQuery(pkg.FirstSubmitted))
		printInfoValue(writer, "Last Modified", formatTimeQuery(pkg.LastModified))

		if pkg.OutOfDate != 0 {
			printInfoValue(writer, "Out-of-date", formatTimeQuery(pkg.OutOfDate))
		} else {
			printInfoValue(writer, "Out-of-date", "No")
		}

		printInfoValue(writer, "ID", fmt.Sprintf("%d", pkg.ID))
		printInfoValue(writer, "Package Base ID", fmt.Sprintf("%d", pkg.PackageBaseID))
		printInfoValue(writer, "Package Base", pkg.PackageBase)
		printInfoValue(writer, "Snapshot URL", aurURL+pkg.URLPath)
	}

	fmt.Fprintln(writer)
}