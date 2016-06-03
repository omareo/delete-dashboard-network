# delete-dashboard-network
Pulls list of networks for org using Meraki dashboard API. Type in the name of a network and watch it disappear.

## Getting this code to run
Please contact Meraki Support to enable API access and retrieve your org ID and shard ID. 
Once API is enabled you will be able to generate an API access key from the profile page.

You will need to plug this information in as environment variables in your OS.

If you are on Mac add to your .bash_profile

export DASHBOARD_API_KEY="myAPIkey"

export DASHBOARD_API_SHARD_ID="myShardID"

export DASHBOARD_API_ORG_ID="myOrgID"

Once these variables are in place can run "go install" from inside the "delete-dashboard-network" directory which will compile a binary.
go to your "bin" directory and execute the binary to run the program.
