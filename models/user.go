package models

import (
	"entgo.io/ent/dialect/entsql"
	"math/rand"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schemas definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {

	// definitly gonna forget to change this every update, I will try to keep this as updated as possible tho
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("last_login").Optional().Default(func() time.Time { return time.Time{} }),
		field.String("registration_ip").Default(""),
		field.String("avatar").Optional().Nillable(), // This is the name that the cdn stores it as

		field.Int("uid").Default(0),

		//authentication
		field.String("username").Unique(),
		field.String("email").Unique(),
		field.String("password"), //Sensitive(),
		field.String("key").NotEmpty(),

		// bad schemas design :pain: Will definitly migrate to a seperate settings table soon:tm: // theres also gonna be a lot of .Sensitive being commented out
		// this is to make it show up in the json file
		//settings // they have a storagekey of settings_ because there wasnt a prefix before, and this makes it so the data isnt lost
		field.Int("upload_settings_url_length").Default(10).StorageKey("settings_url_length"),

		field.JSON("upload_settings_embeds", []Embed{}).Default([]Embed{
			{
				HeaderText:  "File uploaded using nest.rip",
				HeaderURL:   "https://nest.rip",
				AuthorText:  "{{user}}",
				AuthorURL:   "{{user_url}}",
				Title:       "{{filename}}",
				Description: "This is one of the {{uploads}} files I have uploaded",
				Color:       "random",
				Enabled:     func() *bool { b := true; return &b }(),
			},
		}).StorageKey("settings_embeds"),
		field.Enum("upload_settings_url_type").Values("NORMAL", "EMOJI", "INVISIBLE", "RAW").Default("NORMAL").StorageKey("settings_url_type"),
		field.Bool("upload_settings_url_show_extension").Default(false).StructTag(`json:"upload_settings_url_show_extension"`).StorageKey("settings_url_show_extension"),
		field.Bool("upload_settings_should_embed").Default(true).StructTag(`json:"upload_settings_should_embed"`).StorageKey("settings_should_embed"),
		field.JSON("upload_settings_url_domain_settings", []DomainSetting{}).Default([]DomainSetting{
			{
				Domain:    "nest.rip",
				Subdomain: "",
			},
		}).StorageKey("settings_url_domain_settings"),
		// The age of files to delete in seconds (set to 30 days by default for free users)
		field.Uint64("upload_settings_autowipe").Default(0).StructTag(`json:"upload_settings_autowipe"`),
		field.Bool("upload_settings_exploding").Default(false).StructTag(`json:"upload_settings_exploding"`),

		// continue
		field.Int("invites").Default(0).StructTag(`json:"invites"`),

		//Profile settings
		field.String("profile_bio").Default("").StructTag(`json:"profile_bio"`),
		field.Bool("profile_private").Default(false).StructTag(`json:"profile_private"`),

		field.JSON("blacklisted", Blacklisted{}).Default(Blacklisted{
			IsBlacklisted: false,
			Reason:        "",
		}).Optional().StructTag(`json:"blacklisted"`),

		//Discord info
		field.String("discord_id").Default("").StructTag(`json:"discord_id"`),
		field.String("discord_avatar").Default("").StructTag(`json:"discord_avatar"`),
		field.String("discord_refresh_token").Default(""), //Sensitive(),
		field.String("discord_user").Default("").StructTag(`json:"discord_user"`),

		// This is there, so I can invalidate my jwt tokens, on password change for example
		field.Int64("tokenversion").Default(int64(rand.Uint32())), //StructTag(`json:"-"`),

		field.Enum("rank").Values("FREE", "NORMAL", "PREMIUM", "MODERATOR", "ADMIN", "INVITED", "OWNER").Default("INVITED").StructTag(`json:"rank"`),
		field.Int64("premium_expires_at").Default(-1),                                                     //StructTag(`json:"-"`),
		field.Time("last_testimonial_update").Optional().Default(func() time.Time { return time.Time{} }), //StructTag(`json:"-"`),
		field.Time("last_username_update").Optional().Nillable().StructTag(`json:"last_username_update"`),
		field.Time("last_password_update").Optional().Nillable().StructTag(`json:"last_password_update"`),

		field.Strings("past_usernames").Optional().Default([]string{}),

		field.Bool("email_verified").Default(true).StructTag(`json:"emailVerified"`),
		field.String("email_verification_key").Optional(), //.Sensitive(),
		field.Time("last_email_update").Optional().Nillable().Default(func() time.Time { return time.Time{} }).StructTag(`json:"last_email_update"`),

		field.String("password_reset_token").Optional().Nillable(), //.Sensitive(), removed to make it actually visible in the json

		// if the account will be deleted
		field.Bool("deleting").Default(false).StructTag(`json:"deleting"`),
		// this is the date the user requested their account to be deleted
		field.Time("requested_deletion_at").Optional().Nillable().Default(func() time.Time { return time.Time{} }).StructTag(`json:"requested_deletion_at"`),

		field.Bool("requesting_data").Default(false).StructTag(`json:"requesting_data"`),
		field.Time("last_data_request").Optional().Nillable().Default(func() time.Time { return time.Time{} }).StructTag(`json:"last_data_request"`),

		field.Int("strikes").Default(0).StructTag(`json:"strikes"`),

		field.Bool("disabled").Default(false).StructTag(`json:"disabled"`),
		field.Bool("twofa_enabled").Default(false).StructTag(`json:"twofa_enabled"`),
		field.String("twofa_secret").Optional().Sensitive(),
		field.Strings("twofa_recovery_codes").Optional(),

		field.Bool("allow_beta").Default(false).StructTag(`json:"allow_beta"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Why not files? There is a seperate service that automatically deletes the files when the user is deleted.
		edge.From("files", File.Type).Ref("uploader"),
		edge.From("motds", MOTD.Type).Ref("creator").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("invite_list", Invite.Type).Ref("creator").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("domains", Domain.Type).Ref("creator").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("usable_domains", Domain.Type).Ref("usable_by").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("invited_users", User.Type).From("invited_by").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("testimonial", Testimonial.Type).Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		edge.To("data_requests", DataRequest.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

// Indexes of the User
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("key"),
		index.Fields("email"),
	}
}

type Embed struct {
	// sitename
	HeaderText string `json:"headerText"`
	HeaderURL  string `json:"headerURL"`

	// author
	AuthorText string `json:"authorText"`
	AuthorURL  string `json:"authorURL"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`

	Enabled *bool `json:"enabled"`
}

type DomainSetting struct {
	Domain     string  `json:"domain"`
	Subdomain  string  `json:"subdomain"`
	FilePrefix *string `json:"filePrefix,omitempty"`
}

type Blacklisted struct {
	IsBlacklisted bool   `json:"isBlacklisted"`
	Reason        string `json:"reason"`
}
