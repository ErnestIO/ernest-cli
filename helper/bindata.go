// Code generated by go-bindata.
// sources:
// lang/en.yml
// DO NOT EDIT!

package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _langEnYml = []byte(`en:
  info:
    usage: "Displays system-wide information"
    args: " "
    description: |
      Info will display your current login information

      Example:
      $ ernest info
      Target:      http://127.0.0.1:8081
      User:        usr
      CLI Version: 2.2.0
  user:
    usage: "User related subcommands"
    list:
      usage: "List available users."
      args: " "
      description: |
        List available users.

      Example:
          $ ernest user list
    admin:
      usage: "Add or remove admin users"
      add:
        usage: "Adds a specific user as ernest admin"
        description: |
          Adds a specific user as ernest admin

        Example:
            $ ernest user admin add john
      rm:
        usage: "Removes a specific user as ernest admin"
        description: |
          Removes a specific user as ernest admin

        Example:
            $ ernest user admin rm john
    create:
      usage: "Create a new user."
      args: "<username> <password>"
      description: |
        Create a new user on the targeted instance of Ernest.
        Example:
          $ ernest user create <username> <password>
          You can also add an email to the user with the flag --email

        Example:
          $ ernest user create --email username@example.com <username> <password>
    change_password:
      usage: "Change password of available users"
      description: |
        Change password of available users.

        Example:
          $ ernest user change-password
        or changing a change-password by being admin:
          $ ernest user change-password --user <username> --current-password <current-password> --password <new-password>
    disable:
      usage: "Disable available users."
      args: "<username>"
      description: |
        Disable available users.

        Example:
          $ ernest user disable <user-name>
    info:
      usage: "Displays information about the specified user (current user by default)."
      description: |
        Example:
          $ ernest user info
          $ ernest user info --user <user-name>
    enable-mfa:
      usage: "Enable Multi-Factor Authentication."
      args: "[--user-name]"
      description: |
        Enables Multi-Factor Authentication for a user.

        Example:
          $ ernest user enable-mfa
    disable-mfa:
      usage: "Disable Multi-Factor Authentication."
      args: "[--user-name]"
      description: |
        Disable Multi-Factor Authentication for a user.

        Example:
          $ ernest user disable-mfa
    reset-mfa:
      usage: "Reset Multi-Factor Authentication."
      args: "[--user-name]"
      description: |
        Generates a new Multi-Factor Authentication token for a user.

        Example:
          $ ernest user reset-mfa
  aws:
    create:
      usage: "Create a new aws project."
      description: |
        Create a new AWS project on the targeted instance of Ernest.
        Example:
          $ ernest project create aws --region us-west-2 --access_key_id AKIAIOSFODNN7EXAMPLE --secret_access_key wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY my_project
        Template example:
          $ ernest project create aws --template myproject.yml myproject
        Where myproject.yaml will look like:
            ---
            fake: true
            access_key_id : AKIAIOSFODNN7EXAMPLE
            secret_access_key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
            region: us-west-2
      args: "<project-name>"
    update:
      usage: "Updates the specified AWS project."
      args: "<project-name>"
      description: |
        Updates the specified AWS project.
      Example:
          $ ernest project update aws --access_key_id AKIAIOSFODNN7EXAMPLE --secret_access_key wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY my_project
  azure:
    create:
      usage: "Create a new azure project."
      args: "<project-name>"
      description: |
        Create a new Azure project on the targeted instance of Ernest.

        Example:
          $ ernest project create azure --region westus --subscription_id SUBSCRIPTION --client_id USER --client_secret PASSWORD --tenant_id TENANT --environment public my_project

        Template example:
          $ ernest project create azure --template myproject.yml myproject
        Where myproject.yaml will look like:
          ---
          fake: true
          region: westus
          subscription_id: SUBSCRIPTION
          client_id: USER
          client_secret: PASSWORD
          tenant_id: TENANT
          environment: public
    update:
      usage: "Updates the specified Azure project."
      args: "<project-name>"
      description: |
        Updates the specified Azure project.

        Example:
          $ ernest project update azure --subscription_id SUBSCRIPTION --client_id USER --client_secret PASSWORD --tenant_id TENANT --environment public my_project
  docs:
    usage: "Open docs in the default browser."
    args: ""
    description: |
      Open docs in the default browser.

      Example:
        $ ernest docs
  envs:
    list:
      usage: "List available environments."
      args: " "
      description: |
        List available environments and shows its most relevant information.

        Example:
          $ ernest environment list
    update:
      usage: "Creates an empty environment based on a specific project"
      args: "<project> <environment>"
      description: |
        You must be logged in to execute this command.

        Examples:
          $ ernest env update --credentials project.yml my_project my_environment
    create:
      usage: "Creates an empty environment based on a specific project"
      args: "<project> <environment>"
      description: |
        You must be logged in to execute this command.

        Examples:
          $ ernest env create my_project my_environment
          $ ernest env create --credentials project.yml my_project my_environment
    apply:
      usage: "Builds or changes infrastructure."
      args: "<file.yml>"
      description: |
        Sends an environment YAML description file to Ernest to be executed.
        You must be logged in to execute this command.

        If the file is not provided, ernest.yml will be used by default.

        Examples:
          $ ernest env apply myenvironment.yml
          $ ernest env apply --dry myenvironment.yml
    destroy:
      usage: "Destroy an environment."
      args: "<project> <environment_name>"
      description: |
        Destroys an environment by name.

        Example:
          $ ernest env delete <my_project> <my_environment>
    history:
      usage: "Shows the history of an environment, a list of builds"
      args: "ernest-cli env history <my_project> <my_env>"
      description: |
        Shows the history of an environment, a list of builds and its status and basic information.

        Example:
          $ ernest env history <my_project> <my_env>
    reset:
      usage: "Reset an in progress environment."
      args: "<env_name>"
      description: |
        Reseting an environment creation may cause problems, please make sure you know what are you doing.

        Example:
          $ ernest env reset <my_env>
    revert:
      usage: "Reverts an environment to a previous state"
      args: "<project> <env_name> <build_id>"
      description: |
        Reverts an environment to a previous known state using a build ID from 'ernest env history'.

        Example:
          $ ernest env revert <project> <env_name> <build_id>
          $ ernest env revert --dry <project> <env_name> <build_id>
    definition:
      usage: "Show the current definition of an environment by its name"
      args: "<project_name> <env_name>"
      description: |
        Show the current definition of an environment by its name getting the definition about the build.

        Example:
          $ ernest env definition <my_project> <my_env>
    info:
      usage: "$ ernest env info <my_env> --build <specific build>"
      args: "<project_name> <env_name>"
      description: |
        Will show detailed information of the last build of a specified environment.
        In case you specify --build option you will be able to output the detailed information of specific build of an environment.

        Examples:
          $ ernest env definition <my_project> <my_env>
          $ ernest env definition <my_project> <my_env> --build build1
    sync:
      usage: "$ ernest env sync <my_project> <my_env>"
      args: "<project_name> <env_name>"
      description: |
        Will sync ernest's environment state from a provider.
        Any changes detected can then be resolved using the 'resolve' command.

        Examples:
          $ ernest env sync <my_project> <my_env>
    resolve:
      usage: "$ ernest env resolve --[accept|reject|ignore] <my_project> <my_env>"
      args: "<project_name> <env_name>"
      description: |
        Provides the ability to manage changes detected by a sync.
        Options:
          accept changes from provider. Ernests internal state will be updated.
          reject changes from provider. Ernest will create a build to restore and overwrite any changes on the provider.
          ignore changes from provider. Ernest will disgard the results of the last sync.

        Examples:
          $ ernest env resolve --accept <my_project> <my_env>
          $ ernest env resolve --reject <my_project> <my_env>
          $ ernest env resolve --ignore <my_project> <my_env>
    review:
      usage: "$ ernest env review --[accept|reject] <my_project> <my_env>"
      args: "<project_name> <env_name>"
      description: |
        Provides the ability to review submitted builds. Running without any flags will show the diff of the submitted build with the prior environment state.
        Options:
          accept submitted build. This will trigger an environment build
          reject submitted build. The build will be rejected.

        Examples:
          $ ernest env review <my_project> <my_env>
          $ ernest env review --accept <my_project> <my_env>
          $ ernest env review --reject <my_project> <my_env>
    diff:
      usage: "$ ernest env diff <project_name> <env_name> <build_a> <build_b>"
      args: "<env_aname> <build_a> <build_b>"
      description: |
        Will display the diff between two different builds

        Examples:
          $ ernest env diff <my_project> <my_env> 1 2
    import:
      usage: "$ ernest env import <my_project> <my_env>"
      args: "<env_name>"
      description : |
        Will import the environment <my_env> from project <project_name>

        Examples:
          $ ernest env import my_project my_env
  log:
    usage: "Inline display of ernest logs."
    args: " "
    description: |
      Display ernest server logs inline

      Example:
        $ ernest log
        $ ernest log --raw
  login:
    usage: "Login with your Ernest credentials."
    args: " "
    description: |
      Logs an user into Ernest instance.

      Example:
        $ ernest login

      It can also be used without asking the username and password.

      Example:
        $ ernest login --user <user> --password <password>
  logout:
    usage: "Clear local authentication credentials."
    args: " "
    description: |
      Logs out an user from Ernest instance.

      Example:
        $ ernest logout
  monitor:
    usage: "Monitor an environment creation."
    args: "<project_name> <env_name>"
    description: |
      Monitors an environment while it is being built by its name.

      Example:
        $ ernest monitor <my_project> <my_env>
  notification:
    list:
      usage: "List available notifications."
      args: " "
      description: |
        List available notifications.

        Example:
          $ ernest notification list
    delete:
      usage: "Deletes an existing notify."
      args: "<notify_name>"
      description: |
        Deletes an existing notify on the targeted instance of Ernest.

        Example:
          $ ernest notify delete <notify_name>


        Example:
        $ ernest notify delete my_notify
    update:
      usage: "Update a new notify."
      args: "<notify_name> <notify_config>"
      description: |
        Update an existing notify on the targeted instance of Ernest.

        Example:
          $ ernest notify update <notify_name> <provider-details>


        Example:
        $ ernest notify update my_notify '{"url":"https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"}'
    service:
      add:
        usage: "Add environment to an existing notify."
        args: "<notification_name> <project_name> [<env_name>]"
        description: |
          Adds a environment to an existing notify.

          Example:
            $ ernest notify add <notify_name> <project_name> <environment_name>


          Example:
          $ ernest notify add my_notify my_project
          $ ernest notify add my_notify my_project my_env
      rm:
        usage: "Removes an environment to an existing notify."
        args: "<notify_name> <project_name> [<env_name>]"
        description: |
          Removes an environment to an existing notify.

          Example:
            $ ernest notify remove <notify_name> <project_name> <env_name>


          Example:
          $ ernest notify remove my_notify my_project
          $ ernest notify remove my_notify my_project my_env
    create:
      usage: "Create a new notify."
      args: "<notify_name> <notify_type> <notify_config>"
      description: |
        Create a new notify on the targeted instance of Ernest.

        Example:
          $ ernest notify create <notify_name> <provider_type> <provider-details>


        Example:
        $ ernest notify create my_notify slack '{"url":"https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"}'
  logger:
    list:
      usage: "Lists active loggers."
      args: " "
      description: |
        List active loggers.

        Example:
          $ ernest preferences logger list
    set:
      usage: "Creates / updates a logger based on its type."
      args: " "
      description: |
        Creates / updates a logger based on its types.

        Example:
          $ ernest preferences logger add basic --logfile /tmp/ernest.log
          $ ernest preferences logger add logstash --hostname 10.50.1.1 --port 5000 --timeout 50000
          $ ernest preferences logger add rollbar --token MY_ROLLBAR_TOKEN
    del:
      usage: "Deletes a logger based on its type."
      args: " "
      description: |
        Deletes a logger based on its types.

        Example:
          $ ernest preferences logger delete basic
  project:
    list:
      usage: "List available projects."
      args: " "
      description: |
        List available projects.

        Example:
          $ ernest project list
    info:
      usage: "Project information"
      args: " "
      description: |
        Display specific project information.

        Example:
          $ ernest project info <my_project>
  roles:
    set:
      usage: "ernest role set -u john -r owner -p project"
      args: " "
      description: |
        Set permissions for a user on a specific resource

        Example:
          $ ernest roles set -u john -r owner -p my_project
          $ ernest roles set -u john -r reader -p my_project -e my_environment
    unset:
      usage: "ernest role unset -u john -r owner -p my_project"
      args: " "
      description: |
        Set permissions for a user on a specific resource

        Example:
          $ ernest roles set -u john -r owner -p my_project
          $ ernest roles set -u john -r reader -p my_project -e my_environment
  target:
    usage: "Configure Ernest target instance."
    args: "<ernest_url>"
    description: |
      Sets up ernest instance target.

      Example:
        $ ernest target https://myernest.com
  usage:
    usage: "Exports an usage report to the current folder"
    args: " "
    description: |
      Example:
        $ ernest usage --from 2017-01-01 --to 2017-02-01 --output=report.log
        A file named report.log has been exported to the current folder

      Example 2:
        $ ernest usage > myreport.log
  vcloud:
    create:
      usage: "Create a new vcloud project."
      args: "<project-name>"
      description: |
        Create a new vcloud project on the targeted instance of Ernest.

        Example:
          $ ernest project create vcloud --user username --password xxxx --org MY-ORG-NAME --vse-url http://vse.url --vcloud-url https://myernest.com --public-network MY-PUBLIC-NETWORK myproject

        Template example:
          $ ernest project create vcloud --template myproject.yml myproject
          Where myproject.yaml will look like:
            ---
            fake: true
            org: org
            password: pwd
            public-network: MY-NETWORK
            user: bla
            vcloud-url: "http://ss.com"
            vse-url: "http://ss.com"
    delete:
      usage: "Deletes the specified project."
      args: "<project-name>"
      description: |
        Deletes the name specified project.

        Example:
          $ ernest project delete my_project
    update:
      usage: "Updates the specified VCloud project."
      args: "<project-name>"
      description: |
        Updates the specified VCloud project.

        Example:
          $ ernest project update vcloud --user <me> --org <org> --password <secret> my_project
`)

func langEnYmlBytes() ([]byte, error) {
	return _langEnYml, nil
}

func langEnYml() (*asset, error) {
	bytes, err := langEnYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lang/en.yml", size: 17517, mode: os.FileMode(420), modTime: time.Unix(1511351654, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"lang/en.yml": langEnYml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"lang": &bintree{nil, map[string]*bintree{
		"en.yml": &bintree{langEnYml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

