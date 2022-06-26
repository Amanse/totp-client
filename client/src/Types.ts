export interface Response {
    issuers: Issuer
}

interface Issuer {
    [name: string]: Accounts
}

export interface Accounts {
    accounts: Account
}

interface Account {
    [name: string]: string
}
