// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package hyperone

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/hcl2template"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string                      `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string                      `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion         *string                      `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug               *bool                        `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool                        `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string                      `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string            `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string                     `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Type                      *string                      `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect        *string                      `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                   *string                      `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                   *int                         `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername               *string                      `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword               *string                      `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName            *string                      `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string                      `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType   *string                      `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits   *int                         `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                []string                     `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys    *bool                        `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos               []string                     `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile         *string                      `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile        *string                      `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                    *bool                        `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                *string                      `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout            *string                      `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth              *bool                        `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool                        `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int                         `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost            *string                      `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort            *int                         `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool                        `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string                      `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword        *string                      `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive     *bool                        `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string                      `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile *string                      `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod     *string                      `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost              *string                      `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort              *int                         `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername          *string                      `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword          *string                      `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string                      `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string                      `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string                     `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string                     `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey              []byte                       `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey             []byte                       `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                 *string                      `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword             *string                      `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                 *string                      `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy              *bool                        `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                 *int                         `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout              *string                      `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL               *bool                        `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure             *bool                        `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM              *bool                        `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	APIURL                    *string                      `mapstructure:"api_url" required:"false" cty:"api_url" hcl:"api_url"`
	Token                     *string                      `mapstructure:"token" required:"true" cty:"token" hcl:"token"`
	Project                   *string                      `mapstructure:"project" required:"true" cty:"project" hcl:"project"`
	TokenLogin                *string                      `mapstructure:"token_login" required:"false" cty:"token_login" hcl:"token_login"`
	StateTimeout              *string                      `mapstructure:"state_timeout" required:"false" cty:"state_timeout" hcl:"state_timeout"`
	SourceImage               *string                      `mapstructure:"source_image" required:"true" cty:"source_image" hcl:"source_image"`
	ImageName                 *string                      `mapstructure:"image_name" required:"false" cty:"image_name" hcl:"image_name"`
	ImageDescription          *string                      `mapstructure:"image_description" required:"false" cty:"image_description" hcl:"image_description"`
	ImageTags                 map[string]string            `mapstructure:"image_tags" required:"false" cty:"image_tags" hcl:"image_tags"`
	ImageTag                  []hcl2template.FlatKeyValue  `mapstructure:"image_tag" required:"false" cty:"image_tag" hcl:"image_tag"`
	ImageService              *string                      `mapstructure:"image_service" required:"false" cty:"image_service" hcl:"image_service"`
	VmType                    *string                      `mapstructure:"vm_type" required:"true" cty:"vm_type" hcl:"vm_type"`
	VmName                    *string                      `mapstructure:"vm_name" required:"false" cty:"vm_name" hcl:"vm_name"`
	VmTags                    map[string]string            `mapstructure:"vm_tags" required:"false" cty:"vm_tags" hcl:"vm_tags"`
	VmTag                     []hcl2template.FlatNameValue `mapstructure:"vm_tag" required:"false" cty:"vm_tag" hcl:"vm_tag"`
	DiskName                  *string                      `mapstructure:"disk_name" required:"false" cty:"disk_name" hcl:"disk_name"`
	DiskType                  *string                      `mapstructure:"disk_type" required:"false" cty:"disk_type" hcl:"disk_type"`
	DiskSize                  *float32                     `mapstructure:"disk_size" required:"true" cty:"disk_size" hcl:"disk_size"`
	Network                   *string                      `mapstructure:"network" required:"false" cty:"network" hcl:"network"`
	PrivateIP                 *string                      `mapstructure:"private_ip" required:"false" cty:"private_ip" hcl:"private_ip"`
	PublicIP                  *string                      `mapstructure:"public_ip" required:"false" cty:"public_ip" hcl:"public_ip"`
	PublicNetAdpService       *string                      `mapstructure:"public_netadp_service" required:"false" cty:"public_netadp_service" hcl:"public_netadp_service"`
	ChrootDevice              *string                      `mapstructure:"chroot_device" cty:"chroot_device" hcl:"chroot_device"`
	ChrootDisk                *bool                        `mapstructure:"chroot_disk" cty:"chroot_disk" hcl:"chroot_disk"`
	ChrootDiskSize            *float32                     `mapstructure:"chroot_disk_size" cty:"chroot_disk_size" hcl:"chroot_disk_size"`
	ChrootDiskType            *string                      `mapstructure:"chroot_disk_type" cty:"chroot_disk_type" hcl:"chroot_disk_type"`
	ChrootMountPath           *string                      `mapstructure:"chroot_mount_path" cty:"chroot_mount_path" hcl:"chroot_mount_path"`
	ChrootMounts              [][]string                   `mapstructure:"chroot_mounts" cty:"chroot_mounts" hcl:"chroot_mounts"`
	ChrootCopyFiles           []string                     `mapstructure:"chroot_copy_files" cty:"chroot_copy_files" hcl:"chroot_copy_files"`
	ChrootCommandWrapper      *string                      `mapstructure:"chroot_command_wrapper" cty:"chroot_command_wrapper" hcl:"chroot_command_wrapper"`
	MountOptions              []string                     `mapstructure:"mount_options" cty:"mount_options" hcl:"mount_options"`
	MountPartition            *string                      `mapstructure:"mount_partition" cty:"mount_partition" hcl:"mount_partition"`
	PreMountCommands          []string                     `mapstructure:"pre_mount_commands" cty:"pre_mount_commands" hcl:"pre_mount_commands"`
	PostMountCommands         []string                     `mapstructure:"post_mount_commands" cty:"post_mount_commands" hcl:"post_mount_commands"`
	SSHKeys                   []string                     `mapstructure:"ssh_keys" required:"false" cty:"ssh_keys" hcl:"ssh_keys"`
	UserData                  *string                      `mapstructure:"user_data" required:"false" cty:"user_data" hcl:"user_data"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":          &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":      &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":      &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                  &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":  &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":         &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file": &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":               &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"api_url":                      &hcldec.AttrSpec{Name: "api_url", Type: cty.String, Required: false},
		"token":                        &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"project":                      &hcldec.AttrSpec{Name: "project", Type: cty.String, Required: false},
		"token_login":                  &hcldec.AttrSpec{Name: "token_login", Type: cty.String, Required: false},
		"state_timeout":                &hcldec.AttrSpec{Name: "state_timeout", Type: cty.String, Required: false},
		"source_image":                 &hcldec.AttrSpec{Name: "source_image", Type: cty.String, Required: false},
		"image_name":                   &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_description":            &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"image_tags":                   &hcldec.AttrSpec{Name: "image_tags", Type: cty.Map(cty.String), Required: false},
		"image_tag":                    &hcldec.BlockListSpec{TypeName: "image_tag", Nested: hcldec.ObjectSpec((*hcl2template.FlatKeyValue)(nil).HCL2Spec())},
		"image_service":                &hcldec.AttrSpec{Name: "image_service", Type: cty.String, Required: false},
		"vm_type":                      &hcldec.AttrSpec{Name: "vm_type", Type: cty.String, Required: false},
		"vm_name":                      &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"vm_tags":                      &hcldec.AttrSpec{Name: "vm_tags", Type: cty.Map(cty.String), Required: false},
		"vm_tag":                       &hcldec.BlockListSpec{TypeName: "vm_tag", Nested: hcldec.ObjectSpec((*hcl2template.FlatNameValue)(nil).HCL2Spec())},
		"disk_name":                    &hcldec.AttrSpec{Name: "disk_name", Type: cty.String, Required: false},
		"disk_type":                    &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"network":                      &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"private_ip":                   &hcldec.AttrSpec{Name: "private_ip", Type: cty.String, Required: false},
		"public_ip":                    &hcldec.AttrSpec{Name: "public_ip", Type: cty.String, Required: false},
		"public_netadp_service":        &hcldec.AttrSpec{Name: "public_netadp_service", Type: cty.String, Required: false},
		"chroot_device":                &hcldec.AttrSpec{Name: "chroot_device", Type: cty.String, Required: false},
		"chroot_disk":                  &hcldec.AttrSpec{Name: "chroot_disk", Type: cty.Bool, Required: false},
		"chroot_disk_size":             &hcldec.AttrSpec{Name: "chroot_disk_size", Type: cty.Number, Required: false},
		"chroot_disk_type":             &hcldec.AttrSpec{Name: "chroot_disk_type", Type: cty.String, Required: false},
		"chroot_mount_path":            &hcldec.AttrSpec{Name: "chroot_mount_path", Type: cty.String, Required: false},
		"chroot_mounts":                &hcldec.AttrSpec{Name: "chroot_mounts", Type: cty.List(cty.List(cty.String)), Required: false},
		"chroot_copy_files":            &hcldec.AttrSpec{Name: "chroot_copy_files", Type: cty.List(cty.String), Required: false},
		"chroot_command_wrapper":       &hcldec.AttrSpec{Name: "chroot_command_wrapper", Type: cty.String, Required: false},
		"mount_options":                &hcldec.AttrSpec{Name: "mount_options", Type: cty.List(cty.String), Required: false},
		"mount_partition":              &hcldec.AttrSpec{Name: "mount_partition", Type: cty.String, Required: false},
		"pre_mount_commands":           &hcldec.AttrSpec{Name: "pre_mount_commands", Type: cty.List(cty.String), Required: false},
		"post_mount_commands":          &hcldec.AttrSpec{Name: "post_mount_commands", Type: cty.List(cty.String), Required: false},
		"ssh_keys":                     &hcldec.AttrSpec{Name: "ssh_keys", Type: cty.List(cty.String), Required: false},
		"user_data":                    &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
	}
	return s
}
