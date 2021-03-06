/*
** Copyright [2013-2016] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package debian

import (
	"github.com/megamsys/libmegdc/templates"
	u "github.com/megamsys/libmegdc/templates/ubuntu"
	"github.com/megamsys/urknall"
)

var debiansnowflakeinstall *DebianSnowflakeInstall

func init() {
	debiansnowflakeinstall = &DebianSnowflakeInstall{}
	templates.Register("DebianSnowflakeInstall", debiansnowflakeinstall)
}

type DebianSnowflakeInstall struct{}

func (tpl *DebianSnowflakeInstall) Render(p urknall.Package) {
	p.AddTemplate("snowflake", &DebianSnowflakeInstallTemplate{})
}

func (tpl *DebianSnowflakeInstall) Options(t *templates.Template) {
}

func (tpl *DebianSnowflakeInstall) Run(target urknall.Target,inputs []string) error {
	return urknall.Run(target, &DebianSnowflakeInstall{},inputs)
}

type DebianSnowflakeInstallTemplate struct{}

func (m *DebianSnowflakeInstallTemplate) Render(pkg urknall.Package) {

	pkg.AddCommands("repository",
		u.Shell("echo 'deb [arch=amd64] "+DefaultMegamRepo+"' > "+ListFilePath),
		u.UpdatePackagesOmitError(),
	)

	pkg.AddCommands("verticesnowflake",
		u.InstallPackages("verticesnowflake"),
	)
}
