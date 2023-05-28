package mysql

import (
	"time"
)

// Workspace model.
type Workspace struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	Description string
	Storage     WorkspaceStorage `gorm:"serializer:json"`
	CreateTime  time.Time
	UpdateTime  time.Time
}

// WorkspaceStorage ...
type WorkspaceStorage struct {
	NFS *NFSWorkspaceStorage `json:"nfs,omitempty"`
}

// NFSWorkspaceStorage ...
type NFSWorkspaceStorage struct {
	MountPath string `json:"mountPath"`
}

func (w *Workspace) TableName() string {
	return "workspace"
}
