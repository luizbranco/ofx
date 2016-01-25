package ofx

import "encoding/xml"

type Banking struct {
	BankingResponse []BankingResponse `xml:"STMTTRNRS"`
}

type BankingResponse struct {
	TransactionUID TransactionUID `xml:"TRNUID"`
	ClientCookie   string         `xml:"CLTCOOKIE,omitempty"`
	Status         Status         `xml:"STATUS"`
}

type BankStatementResponse struct {
	XMLName              xml.Name             `xml:"STMTRS"`
	CurrencyDefault      CurrencySymbol       `xml:"CURDEF"`
	BankingAccount       BankingAccount       `xml:"BANKACCTFROM"`
	BankTransactionsList BankTransactionsList `xml:"BANKTRANLIST,omitempty"`
	LedgerBalance        Balance              `xml:"LEDGERBAL"`
	AvailableBalance     Balance              `xml:"AVAILBAL,omitempty"`
	MarketingInfo        string               `xml:"MKTGINFO,omitempty"`
}

type BankingAccount struct {
	BankID      string      `xml:"BANKID"`
	BranchID    string      `xml:"BRANCHID,omitempty"`
	ID          string      `xml:"ACCTID"`
	AccountType AccountType `xml:"ACCTTYPE"`
	Key         string      `xml:"ACCTKEY,omitempty"`
}
