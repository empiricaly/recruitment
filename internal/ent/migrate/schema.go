// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:        "admins",
		Columns:     AdminsColumns,
		PrimaryKey:  []*schema.Column{AdminsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DataColumns holds the columns for the "data" table.
	DataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "key", Type: field.TypeString, Size: 1024},
		{Name: "val", Type: field.TypeString, Size: 4096},
		{Name: "index", Type: field.TypeInt},
		{Name: "current", Type: field.TypeBool, Default: true},
		{Name: "version", Type: field.TypeInt},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "participant_data", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// DataTable holds the schema information for the "data" table.
	DataTable = &schema.Table{
		Name:       "data",
		Columns:    DataColumns,
		PrimaryKey: []*schema.Column{DataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "data_participants_data",
				Columns: []*schema.Column{DataColumns[9]},

				RefColumns: []*schema.Column{ParticipantsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ParticipantsColumns holds the columns for the "participants" table.
	ParticipantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "mturk_worker_id", Type: field.TypeString, Nullable: true},
		{Name: "uninitialized", Type: field.TypeBool, Nullable: true},
		{Name: "step_run_created_participants", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// ParticipantsTable holds the schema information for the "participants" table.
	ParticipantsTable = &schema.Table{
		Name:       "participants",
		Columns:    ParticipantsColumns,
		PrimaryKey: []*schema.Column{ParticipantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "participants_step_runs_createdParticipants",
				Columns: []*schema.Column{ParticipantsColumns[5]},

				RefColumns: []*schema.Column{StepRunsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ParticipationsColumns holds the columns for the "participations" table.
	ParticipationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "added_participant", Type: field.TypeBool},
		{Name: "mturk_worker_id", Type: field.TypeString},
		{Name: "mturk_assignment_id", Type: field.TypeString},
		{Name: "mturk_hit_id", Type: field.TypeString},
		{Name: "mturk_accepted_at", Type: field.TypeTime, Nullable: true},
		{Name: "mturk_submitted_at", Type: field.TypeTime, Nullable: true},
		{Name: "participant_participations", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "step_run_participations", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// ParticipationsTable holds the schema information for the "participations" table.
	ParticipationsTable = &schema.Table{
		Name:       "participations",
		Columns:    ParticipationsColumns,
		PrimaryKey: []*schema.Column{ParticipationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "participations_participants_participations",
				Columns: []*schema.Column{ParticipationsColumns[9]},

				RefColumns: []*schema.Column{ParticipantsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "participations_step_runs_participations",
				Columns: []*schema.Column{ParticipationsColumns[10]},

				RefColumns: []*schema.Column{StepRunsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "project_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "admin_projects", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "projects_admins_projects",
				Columns: []*schema.Column{ProjectsColumns[5]},

				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProviderIdsColumns holds the columns for the "provider_ids" table.
	ProviderIdsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "mturk_worker_id", Type: field.TypeString},
		{Name: "participant_provider_ids", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// ProviderIdsTable holds the schema information for the "provider_ids" table.
	ProviderIdsTable = &schema.Table{
		Name:       "provider_ids",
		Columns:    ProviderIdsColumns,
		PrimaryKey: []*schema.Column{ProviderIdsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "provider_ids_participants_providerIDs",
				Columns: []*schema.Column{ProviderIdsColumns[4]},

				RefColumns: []*schema.Column{ParticipantsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RunsColumns holds the columns for the "runs" table.
	RunsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"CREATED", "RUNNING", "PAUSED", "DONE", "TERMINATED", "FAILED"}},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "ended_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "start_at", Type: field.TypeTime, Nullable: true},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "project_runs", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "run_current_step", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// RunsTable holds the schema information for the "runs" table.
	RunsTable = &schema.Table{
		Name:       "runs",
		Columns:    RunsColumns,
		PrimaryKey: []*schema.Column{RunsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "runs_projects_runs",
				Columns: []*schema.Column{RunsColumns[9]},

				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "runs_step_runs_currentStep",
				Columns: []*schema.Column{RunsColumns[10]},

				RefColumns: []*schema.Column{StepRunsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StepsColumns holds the columns for the "steps" table.
	StepsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"MTURK_HIT", "MTURK_MESSAGE", "PARTICIPANT_FILTER", "WAIT"}},
		{Name: "index", Type: field.TypeInt},
		{Name: "duration", Type: field.TypeInt},
		{Name: "msg_args", Type: field.TypeJSON, Nullable: true},
		{Name: "hit_args", Type: field.TypeJSON, Nullable: true},
		{Name: "filter_args", Type: field.TypeJSON, Nullable: true},
		{Name: "step_run_step", Type: field.TypeString, Unique: true, Nullable: true, Size: 20},
		{Name: "template_steps", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// StepsTable holds the schema information for the "steps" table.
	StepsTable = &schema.Table{
		Name:       "steps",
		Columns:    StepsColumns,
		PrimaryKey: []*schema.Column{StepsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "steps_step_runs_step",
				Columns: []*schema.Column{StepsColumns[9]},

				RefColumns: []*schema.Column{StepRunsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "steps_templates_steps",
				Columns: []*schema.Column{StepsColumns[10]},

				RefColumns: []*schema.Column{TemplatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StepRunsColumns holds the columns for the "step_runs" table.
	StepRunsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"CREATED", "RUNNING", "PAUSED", "DONE", "TERMINATED", "FAILED"}},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "ended_at", Type: field.TypeTime, Nullable: true},
		{Name: "index", Type: field.TypeInt},
		{Name: "participants_count", Type: field.TypeInt},
		{Name: "hit_id", Type: field.TypeString, Nullable: true},
		{Name: "url_token", Type: field.TypeString, Unique: true},
		{Name: "run_steps", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// StepRunsTable holds the schema information for the "step_runs" table.
	StepRunsTable = &schema.Table{
		Name:       "step_runs",
		Columns:    StepRunsColumns,
		PrimaryKey: []*schema.Column{StepRunsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "step_runs_runs_steps",
				Columns: []*schema.Column{StepRunsColumns[10]},

				RefColumns: []*schema.Column{RunsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TemplatesColumns holds the columns for the "templates" table.
	TemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "selection_type", Type: field.TypeEnum, Enums: []string{"INTERNAL_DB", "MTURK_QUALIFICATIONS"}},
		{Name: "participant_count", Type: field.TypeInt},
		{Name: "internal_criteria", Type: field.TypeJSON},
		{Name: "mturk_criteria", Type: field.TypeJSON},
		{Name: "adult", Type: field.TypeBool},
		{Name: "sandbox", Type: field.TypeBool},
		{Name: "admin_templates", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "project_templates", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "run_template", Type: field.TypeString, Unique: true, Nullable: true, Size: 20},
	}
	// TemplatesTable holds the schema information for the "templates" table.
	TemplatesTable = &schema.Table{
		Name:       "templates",
		Columns:    TemplatesColumns,
		PrimaryKey: []*schema.Column{TemplatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "templates_admins_templates",
				Columns: []*schema.Column{TemplatesColumns[10]},

				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "templates_projects_templates",
				Columns: []*schema.Column{TemplatesColumns[11]},

				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "templates_runs_template",
				Columns: []*schema.Column{TemplatesColumns[12]},

				RefColumns: []*schema.Column{RunsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AdminImportedParticipantsColumns holds the columns for the "admin_importedParticipants" table.
	AdminImportedParticipantsColumns = []*schema.Column{
		{Name: "admin_id", Type: field.TypeString, Size: 20},
		{Name: "participant_id", Type: field.TypeString, Size: 20},
	}
	// AdminImportedParticipantsTable holds the schema information for the "admin_importedParticipants" table.
	AdminImportedParticipantsTable = &schema.Table{
		Name:       "admin_importedParticipants",
		Columns:    AdminImportedParticipantsColumns,
		PrimaryKey: []*schema.Column{AdminImportedParticipantsColumns[0], AdminImportedParticipantsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "admin_importedParticipants_admin_id",
				Columns: []*schema.Column{AdminImportedParticipantsColumns[0]},

				RefColumns: []*schema.Column{AdminsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "admin_importedParticipants_participant_id",
				Columns: []*schema.Column{AdminImportedParticipantsColumns[1]},

				RefColumns: []*schema.Column{ParticipantsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProjectParticipantsColumns holds the columns for the "project_participants" table.
	ProjectParticipantsColumns = []*schema.Column{
		{Name: "project_id", Type: field.TypeString, Size: 20},
		{Name: "participant_id", Type: field.TypeString, Size: 20},
	}
	// ProjectParticipantsTable holds the schema information for the "project_participants" table.
	ProjectParticipantsTable = &schema.Table{
		Name:       "project_participants",
		Columns:    ProjectParticipantsColumns,
		PrimaryKey: []*schema.Column{ProjectParticipantsColumns[0], ProjectParticipantsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "project_participants_project_id",
				Columns: []*schema.Column{ProjectParticipantsColumns[0]},

				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "project_participants_participant_id",
				Columns: []*schema.Column{ProjectParticipantsColumns[1]},

				RefColumns: []*schema.Column{ParticipantsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// StepRunParticipantsColumns holds the columns for the "step_run_participants" table.
	StepRunParticipantsColumns = []*schema.Column{
		{Name: "step_run_id", Type: field.TypeString, Size: 20},
		{Name: "participant_id", Type: field.TypeString, Size: 20},
	}
	// StepRunParticipantsTable holds the schema information for the "step_run_participants" table.
	StepRunParticipantsTable = &schema.Table{
		Name:       "step_run_participants",
		Columns:    StepRunParticipantsColumns,
		PrimaryKey: []*schema.Column{StepRunParticipantsColumns[0], StepRunParticipantsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "step_run_participants_step_run_id",
				Columns: []*schema.Column{StepRunParticipantsColumns[0]},

				RefColumns: []*schema.Column{StepRunsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "step_run_participants_participant_id",
				Columns: []*schema.Column{StepRunParticipantsColumns[1]},

				RefColumns: []*schema.Column{ParticipantsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		DataTable,
		ParticipantsTable,
		ParticipationsTable,
		ProjectsTable,
		ProviderIdsTable,
		RunsTable,
		StepsTable,
		StepRunsTable,
		TemplatesTable,
		AdminImportedParticipantsTable,
		ProjectParticipantsTable,
		StepRunParticipantsTable,
	}
)

func init() {
	DataTable.ForeignKeys[0].RefTable = ParticipantsTable
	ParticipantsTable.ForeignKeys[0].RefTable = StepRunsTable
	ParticipationsTable.ForeignKeys[0].RefTable = ParticipantsTable
	ParticipationsTable.ForeignKeys[1].RefTable = StepRunsTable
	ProjectsTable.ForeignKeys[0].RefTable = AdminsTable
	ProviderIdsTable.ForeignKeys[0].RefTable = ParticipantsTable
	RunsTable.ForeignKeys[0].RefTable = ProjectsTable
	RunsTable.ForeignKeys[1].RefTable = StepRunsTable
	StepsTable.ForeignKeys[0].RefTable = StepRunsTable
	StepsTable.ForeignKeys[1].RefTable = TemplatesTable
	StepRunsTable.ForeignKeys[0].RefTable = RunsTable
	TemplatesTable.ForeignKeys[0].RefTable = AdminsTable
	TemplatesTable.ForeignKeys[1].RefTable = ProjectsTable
	TemplatesTable.ForeignKeys[2].RefTable = RunsTable
	AdminImportedParticipantsTable.ForeignKeys[0].RefTable = AdminsTable
	AdminImportedParticipantsTable.ForeignKeys[1].RefTable = ParticipantsTable
	ProjectParticipantsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectParticipantsTable.ForeignKeys[1].RefTable = ParticipantsTable
	StepRunParticipantsTable.ForeignKeys[0].RefTable = StepRunsTable
	StepRunParticipantsTable.ForeignKeys[1].RefTable = ParticipantsTable
}
