package config

// SampleCfg is a sample config file used to control how eksup works and integrates with AWS
const SampleCfg = `# This is a sample file to configure and control how eksup works
aws:
  auth:
    # Remove this if using SSO profile
    credentials: true
    # Remove the following lines if using AWS access and secret key credentials
    profile: true
    profileName: sso_profile
  # Set the region where the EKS cluster is located
  region: us-east-1
`
