package structs

// PatronDisplay // Output shown to the user
type PatronDisplay struct {
    Name string `json:"full_name"`
    // assign this somehow (or just do it in frontend?)
    //p.Url = "https://www.patreon.com/user?u=" + p.Id
    Id string `json:"id"`
    ImageUrl string `json:"image_url"`
}

type PatronOutput struct {
    Patrons []PatronDisplay `json:"patrons"`
}

// PatreonTokens // Credentials file
type PatreonTokens struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    TokenType string `json:"token_type"`
    Scope string `json:"scope"`
    ExpiresIn int `json:"expires_in"`
}

// PatreonMembersResponse // Patreon api responses
type PatreonMembersResponse struct {
    Data []PatreonMembersData `json:"data"`
    Included []struct {
        Attributes PatronUsersAttribute `json:"attributes"`
        ID string `json:"id"`
        Type string `json:"type"`
    } `json:"included"`
}

type PatreonMembersData struct {
    Type          string                  `json:"type"`
    ID            string                  `json:"id"` // this is the member's id
    Attributes    PatreonMembersAttribute `json:"attributes"`
    Relationships PatronRelationship      `json:"relationships"`
}

type PatronRelationship struct {
    User PatronRelationshipUser `json:"user"`
}

type PatronRelationshipUser struct {
    Data PatronRelationshipUserData `json:"data"`
    Included []struct {
        Attributes PatreonMembersAttribute `json:"attributes"`
        ID string `json:"id"`
        Type string `json:"type"`
    } `json:"included"`
}

type PatronUsersAttribute struct {
    ImageUrl string `json:"image_url"`
}

type PatronRelationshipUserData struct {
    Id   string `json:"id"`
    Type string `json:"type"`
    Relationships struct {
        Memberships struct {
            Data []struct {
                ID string `json:"id"`
                Type string `json:"type"`
            } `json:"data"`
        } `json:"memberships"`
    } `json:"relationships"`
}

type MemberRelationship struct {
    Type string `json:"type"`
    ID   string `json:"id"` // this is the member's id
}

type PatreonMembersAttribute struct {
    FullName           string `json:"full_name"`
    PatronStatus       string `json:"patron_status"`
    WillPayAmountCents int    `json:"will_pay_amount_cents"`
}

type RelatedLink struct {
    Related string `json:"related"`
}

// WebhookPledge adapted to only include the items I need
type WebhookPledge struct {
    Data struct {
        Type       string `json:"type"`
        ID         string `json:"id"` // this is the member id
        Attributes struct {
            PledgeAmountCents int    `json:"pledge_amount_cents"`
            LastChargeStatus  string `json:"last_charge_status"`
            PatronStatus      string `json:"patron_status"`
        } `json:"attributes"`
        Relationships struct {
            User struct {
                Data PatronRelationshipUserData `json:"data"`
                Links RelatedLink               `json:"links"`
            }  `json:"user"`
        } `json:"relationships"`
    } `json:"data"`
}
