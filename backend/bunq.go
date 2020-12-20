package main

import (
	"errors"

	"github.com/OGKevin/go-bunq/bunq"
)

func GetAccountIDForIBAN(c *bunq.Client, iban string) (int, error) {
	accounts, err := c.AccountService.GetAllMonetaryAccountBank()
	if err != nil {
		return 0, err
	}
	savingAccounts, err := c.AccountService.GetAllMonetaryAccountSaving()
	if err != nil {
		return 0, err
	}

	for _, account := range accounts.Response {
		if account.MonetaryAccountBank.GetIBANPointer().Value == iban {
			return account.MonetaryAccountBank.ID, nil
		}
	}

	for _, account := range savingAccounts.Response {
		if account.MonetaryAccountSaving.GetIBANPointer().Value == iban {
			return account.MonetaryAccountSaving.ID, nil
		}
	}

	return 0, errors.New("No account found for IBAN")
}

func GetAccountPoinerForIBAN(c *bunq.Client, iban string) (*bunq.Pointer, error) {
	accounts, err := c.AccountService.GetAllMonetaryAccountBank()
	if err != nil {
		return nil, err
	}
	savingAccounts, err := c.AccountService.GetAllMonetaryAccountSaving()
	if err != nil {
		return nil, err
	}

	for _, account := range accounts.Response {
		if account.MonetaryAccountBank.GetIBANPointer().Value == iban {
			return account.MonetaryAccountBank.GetIBANPointer(), nil
		}
	}

	for _, account := range savingAccounts.Response {
		if account.MonetaryAccountSaving.GetIBANPointer().Value == iban {
			return account.MonetaryAccountSaving.GetIBANPointer(), nil
		}
	}

	return nil, errors.New("No account found for IBAN")
}
