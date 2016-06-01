# delete-dashboard-network
Pulls list of networks for org using Meraki dashboard API. Type in the name of a network and watch it disappear.

## Getting this code to run
You will need to contact Meraki Support to enable API access. Once it is enabled you will be able to generate an API access key on the profile page. Support will be able to provide the org ID and shard ID.


DASHBOARD_API_KEY
DASHBOARD_API_SHARD_ID
DASHBOARD_API_ORG_ID


If you are on Linux/Mac
export DASHBOARD_API_KEY="myAPIkey
export DASHBOARD_API_SHARD_ID="myShardID"
export DASHBOARD_API_ORG_ID="myOrgID"

Once these variables are in place can run "go install" from inside the "delete-dashboard-network" directory which will compile a binary.
go to your "bin" directory and execute the binary to run the program.
