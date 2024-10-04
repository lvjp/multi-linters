# 🐐 Multiple linter on multiple architectures

## Comparison

| Linter                                                       | multi-linter (5) | mega-linter (110) | super-linter (75) |
| ------------------------------------------------------------ | :--------------: | :---------------: | :---------------: |
| [actionlint][actionlint]                                     |                  |        ✅         |        ✅         |
| [ansible-lint][ansible-lint]                                 |                  |        ✅         |        ✅         |
| [Azure Resource Manager Template Toolkit (arm-ttk)][arm-ttk] |                  |        ✅         |        ✅         |
| [ASL Validator][asl-validator]                               |                  |                   |        ✅         |
| [bandit][bandit]                                             |                  |        ✅         |                   |
| bash-exec                                                    |                  |        ✅         |        ✅         |
| [bicep-linter][bicep-linter]                                 |                  |        ✅         |                   |
| [black][black]                                               |                  |        ✅         |        ✅         |
| [AWS CloudFormation Linter (cfn-lint)][cfn-lint]             |                  |        ✅         |        ✅         |
| [checkmake][checkmake]                                       |                  |        ✅         |                   |
| [checkov][checkov]                                           |                  |        ✅         |        ✅         |
| [checkstyle][checkstyle]                                     |                  |        ✅         |        ✅         |
| [chktex][chktex]                                             |                  |                   |        ✅         |
| [clang-format][clang-format]                                 |                  |        ✅         |        ✅         |
| [clippy][clippy]                                             |                  |        ✅         |        ✅         |
| [clj-kondo][clj-kondo]                                       |                  |        ✅         |        ✅         |
| [cljstyle][cljstyle]                                         |                  |        ✅         |                   |
| [coffeelint][coffeelint]                                     |                  |        ✅         |        ✅         |
| [commitlint][commitlint]                                     |                  |                   |        ✅         |
| [cpplint][cpplint]                                           |                  |        ✅         |        ✅         |
| [csharpier][csharpier]                                       |                  |        ✅         |                   |
| [cspell][cspell]                                             |                  |        ✅         |                   |
| [dart-analyze][dart-analyze]                                 |                  |        ✅         |        ✅         |
| [detekt][detekt]                                             |                  |        ✅         |                   |
| [devskim][devskim]                                           |                  |        ✅         |                   |
| [djlint][djlint]                                             |                  |        ✅         |                   |
| [dotenv-linter][dotenv-linter]                               |                  |        ✅         |        ✅         |
| [dotnet-format][dotnet-format]                               |                  |        ✅         |        ✅         |
| [dustilock][dustilock]                                       |                  |        ✅         |                   |
| [editorconfig-checker][editorconfig-checker]                 |                  |        ✅         |        ✅         |
| [eslint-plugin-json][eslint-plugin-json]                     |                  |                   |        ✅         |
| [eslint-plugin-jsonc][eslint-plugin-jsonc]                   |                  |        ✅         |        ✅         |
| [eslint-plugin-jsx-a11y][eslint-plugin-jsx-a11y]             |                  |        ✅         |        ✅         |
| [eslint-plugin-react][eslint-plugin-react]                   |                  |        ✅         |        ✅         |
| [eslint-typescript][eslint-typescript]                       |                  |        ✅         |        ✅         |
| [ESLint][eslint]                                             |                  |        ✅         |        ✅         |
| [flake8][flake8]                                             |                  |        ✅         |        ✅         |
| [gherkin-lint][gherkin-lint]                                 |                  |        ✅         |        ✅         |
| [Git conflict markers presence in files][git-diff]           |                  |        ✅         |        ✅         |
| [GitLeaks][gitleaks]                                         |        ✅        |        ✅         |        ✅         |
| [golangci-lint][golangci-lint]                               |        ✅        |        ✅         |        ✅         |
| [google-java-format][google-java-format]                     |                  |                   |        ✅         |
| [GoReleaser][goreleaser]                                     |                  |                   |        ✅         |
| [graphql-schema-linter][graphql-schema-linter]               |                  |        ✅         |                   |
| [grype][grype]                                               |                  |        ✅         |                   |
| [Haskell Dockerfile Linter][hadolint]                        |        ✅        |        ✅         |        ✅         |
| [helm][helm]                                                 |                  |        ✅         |                   |
| [HTMLHint][htmlhint]                                         |                  |        ✅         |        ✅         |
| [isort][isort]                                               |                  |        ✅         |        ✅         |
| [jscpd][jscpd]                                               |                  |        ✅         |        ✅         |
| [jsonlint][jsonlint]                                         |                  |        ✅         |                   |
| [kics][kics]                                                 |                  |        ✅         |                   |
| [ktlint][ktlint]                                             |                  |        ✅         |        ✅         |
| [kubeconform][kubeconform]                                   |                  |        ✅         |        ✅         |
| [kubescape][kubescape]                                       |                  |        ✅         |                   |
| [LibXML][libxml]                                             |                  |                   |        ✅         |
| [lightning-flow-scanner][lightning-flow-scanner]             |                  |        ✅         |                   |
| [lintr][lintr]                                               |                  |        ✅         |        ✅         |
| [ls-lint][ls-lint]                                           |                  |        ✅         |                   |
| [luacheck][luacheck]                                         |                  |        ✅         |        ✅         |
| [lychee][lychee]                                             |                  |        ✅         |                   |
| [markdown-link-check][markdown-link-check]                   |                  |        ✅         |                   |
| [markdown-table-formatter][markdown-table-formatter]         |                  |        ✅         |                   |
| [markdownlint][markdownlint]                                 |                  |        ✅         |        ✅         |
| [mypy][mypy]                                                 |                  |        ✅         |                   |
| [npm-groovy-lint][npm-groovy-lint]                           |                  |        ✅         |        ✅         |
| [npm-package-json-lint][npm-package-json-lint]               |                  |        ✅         |                   |
| [perlcritic][perlcritic]                                     |                  |        ✅         |        ✅         |
| [php-cs-fixer][php-cs-fixer]                                 |                  |        ✅         |                   |
| [PHP CodeSniffer][phpcs]                                     |                  |        ✅         |        ✅         |
| [phplint][phplint]                                           |                  |        ✅         |                   |
| [PHPStan][phpstan]                                           |                  |        ✅         |        ✅         |
| [PHP built-in linter][phpstd]                                |                  |                   |        ✅         |
| [pmd][pmd]                                                   |                  |        ✅         |                   |
| [powershell-formatter][powershell-formatter]                 |                  |        ✅         |                   |
| [powershell][powershell]                                     |                  |        ✅         |                   |
| [Prettier][prettier]                                         |                  |        ✅         |        ✅         |
| [proselint][proselint]                                       |                  |        ✅         |                   |
| [protolint][protolint]                                       |                  |        ✅         |        ✅         |
| [PSScriptAnalyzer][ps-script-analyzer]                       |                  |                   |        ✅         |
| [Psalm][psalm]                                               |                  |        ✅         |        ✅         |
| [puppet-lint][puppet-lint]                                   |                  |        ✅         |                   |
| [pyink][pyink]                                               |                  |                   |        ✅         |
| [pylint][pylint]                                             |                  |        ✅         |        ✅         |
| [pyright][pyright]                                           |                  |        ✅         |                   |
| [Raku][raku]                                                 |                  |        ✅         |        ✅         |
| [remark-lint][remark-lint]                                   |                  |        ✅         |                   |
| [renovate-config-validator][renovate-config-validator]       |                  |                   |        ✅         |
| [revive][revive]                                             |                  |        ✅         |      ✅[^1]       |
| [roslynator][roslynator]                                     |                  |        ✅         |                   |
| [rst-lint][rst-lint]                                         |                  |        ✅         |                   |
| [rstcheck][rstcheck]                                         |                  |        ✅         |                   |
| [RuboCop][rubocop]                                           |                  |        ✅         |        ✅         |
| [ruff][ruff]                                                 |                  |        ✅         |        ✅         |
| [Rustfmt][rustfmt]                                           |                  |                   |        ✅         |
| [scalafix][scalafix]                                         |                  |        ✅         |                   |
| [scalafmt][scalafmt]                                         |                  |                   |        ✅         |
| [secretlint][secretlint]                                     |                  |        ✅         |                   |
| [semgrep][semgrep]                                           |                  |        ✅         |                   |
| [sfdx-scanner-apex][sfdx-scanner-apex]                       |                  |        ✅         |                   |
| [sfdx-scanner-aura][sfdx-scanner-aura]                       |                  |        ✅         |                   |
| [sfdx-scanner-lwc][sfdx-scanner-lwc]                         |                  |        ✅         |                   |
| [ShellCheck][shellcheck]                                     |        ✅        |        ✅         |        ✅         |
| [shfmt][shfmt]                                               |        ✅        |        ✅         |        ✅         |
| [snakefmt][snakefmt]                                         |                  |        ✅         |        ✅         |
| [snakemake][snakemake]                                       |                  |        ✅         |        ✅         |
| [spectral][spectral]                                         |                  |        ✅         |        ✅         |
| [sqlfluff][sqlfluff]                                         |                  |        ✅         |        ✅         |
| [standardjs][standardjs]                                     |                  |        ✅         |        ✅         |
| [stylelint][stylelint]                                       |                  |        ✅         |        ✅         |
| [swiftlint][swiftlint]                                       |                  |        ✅         |                   |
| [syft][syft]                                                 |                  |        ✅         |                   |
| [tekton-lint][tekton-lint]                                   |                  |        ✅         |        ✅         |
| [terraform-fmt][terraform-fmt]                               |                  |        ✅         |        ✅         |
| [terragrunt][terragrunt]                                     |                  |        ✅         |        ✅         |
| [terrascan][terrascan]                                       |                  |        ✅         |        ✅         |
| [textlint][textlint]                                         |                  |                   |        ✅         |
| [tflint][tflint]                                             |                  |        ✅         |        ✅         |
| [trivy][trivy]                                               |                  |        ✅         |                   |
| [trufflehog][trufflehog]                                     |                  |        ✅         |                   |
| [tsqllint][tsqllint]                                         |                  |        ✅         |                   |
| [v8r][v8r]                                                   |                  |        ✅         |                   |
| [vale][vale]                                                 |                  |        ✅         |                   |
| [YamlLint][yamllint]                                         |                  |        ✅         |        ✅         |

[^1]: Revive is activated through golangci-lint configuration

[actionlint]: https://github.com/rhysd/actionlint
[ansible-lint]: https://github.com/ansible/ansible-lint
[arm-ttk]: https://github.com/Azure/arm-ttk
[asl-validator]: https://github.com/ChristopheBougere/asl-validator
[bandit]: https://github.com/PyCQA/bandit
[bicep-linter]: https://github.com/Azure/bicep
[black]: https://github.com/psf/black
[cfn-lint]: https://github.com/aws-cloudformation/cfn-lint
[checkmake]: https://github.com/mrtazz/checkmake
[checkov]: https://github.com/bridgecrewio/checkov
[checkstyle]: https://github.com/checkstyle/checkstyle

<!-- markdown-link-check-disable -->
<!-- Disable chktex link checking since it has been down for several days -->

[chktex]: https://localhost/git.savannah.nongnu.org/cgit/chktex.git

<!-- markdown-link-check-enable -->

[clang-format]: https://github.com/llvm/llvm-project
[clippy]: https://github.com/rust-lang/rust-clippy
[clj-kondo]: https://github.com/borkdude/clj-kondo
[cljstyle]: https://github.com/greglook/cljstyle
[coffeelint]: https://github.com/clutchski/coffeelint
[commitlint]: https://github.com/conventional-changelog/commitlint
[cpplint]: https://github.com/cpplint/cpplint
[csharpier]: https://github.com/belav/csharpier
[cspell]: https://github.com/streetsidesoftware/cspell
[dart-analyze]: https://github.com/dart-lang/sdk
[detekt]: https://github.com/detekt/detekt
[devskim]: https://github.com/microsoft/DevSkim
[djlint]: https://github.com/Riverside-Healthcare/djlint
[dotenv-linter]: https://github.com/dotenv-linter/dotenv-linter
[dotnet-format]: https://github.com/dotnet/sdk
[dustilock]: https://github.com/Checkmarx/dustilock
[editorconfig-checker]: https://github.com/editorconfig-checker/editorconfig-checker
[eslint-plugin-json]: https://github.com/azeemba/eslint-plugin-json
[eslint-plugin-jsonc]: https://github.com/ota-meshi/eslint-plugin-jsonc
[eslint-plugin-jsx-a11y]: https://github.com/jsx-eslint/eslint-plugin-react
[eslint-plugin-react]: https://github.com/yannickcr/eslint-plugin-react
[eslint-typescript]: https://github.com/typescript-eslint/typescript-eslint
[eslint]: https://github.com/eslint/eslint
[flake8]: https://github.com/PyCQA/flake8
[gherkin-lint]: https://github.com/vsiakka/gherkin-lint
[git-diff]: https://github.com/git/git
[gitleaks]: https://github.com/gitleaks/gitleaks
[golangci-lint]: https://github.com/golangci/golangci-lint
[google-java-format]: https://github.com/google/google-java-format
[goreleaser]: https://github.com/goreleaser/goreleaser
[graphql-schema-linter]: https://github.com/cjoudrey/graphql-schema-linter
[grype]: https://github.com/anchore/grype
[hadolint]: https://github.com/hadolint/hadolint
[helm]: https://github.com/helm/helm
[htmlhint]: https://github.com/htmlhint/HTMLHint
[isort]: https://github.com/PyCQA/isort
[jscpd]: https://github.com/kucherenko/jscpd
[jsonlint]: https://github.com/prantlf/jsonlint
[kics]: https://github.com/checkmarx/kics
[ktlint]: https://github.com/pinterest/ktlint
[kubeconform]: https://github.com/yannh/kubeconform
[kubescape]: https://github.com/kubescape/kubescape
[libxml]: https://gitlab.gnome.org/GNOME/libxml2
[lightning-flow-scanner]: https://github.com/Lightning-Flow-Scanner/lightning-flow-scanner-sfdx
[lintr]: https://github.com/r-lib/lintr
[ls-lint]: https://github.com/loeffel-io/ls-lint
[luacheck]: https://github.com/luarocks/luacheck
[lychee]: https://github.com/lycheeverse/lychee
[markdown-link-check]: https://github.com/tcort/markdown-link-check
[markdown-table-formatter]: https://github.com/nvuillam/markdown-table-formatter
[markdownlint]: https://github.com/DavidAnson/markdownlint
[mypy]: https://github.com/python/mypy
[npm-groovy-lint]: https://github.com/nvuillam/npm-groovy-lint
[npm-package-json-lint]: https://github.com/tclindner/npm-package-json-lint
[perlcritic]: https://github.com/Perl-Critic/Perl-Critic
[php-cs-fixer]: https://github.com/PHP-CS-Fixer/PHP-CS-Fixer
[phpcs]: https://github.com/PHPCSStandards/PHP_CodeSniffer
[phplint]: https://github.com/overtrue/phplint
[phpstan]: https://github.com/phpstan/phpstan
[phpstd]: https://github.com/php/php-src
[pmd]: https://github.com/pmd/pmd
[powershell-formatter]: https://github.com/PowerShell/PSScriptAnalyzer
[powershell]: https://github.com/PowerShell/PSScriptAnalyzer
[prettier]: https://github.com/prettier/prettier
[proselint]: https://github.com/amperser/proselint
[protolint]: https://github.com/yoheimuta/protolint
[ps-script-analyzer]: https://github.com/PowerShell/Psscriptanalyzer
[psalm]: https://github.com/vimeo/psalm
[puppet-lint]: https://github.com/puppetlabs/puppet-lint
[pyink]: https://github.com/google/pyink
[pylint]: https://github.com/pylint-dev/pylint
[pyright]: https://github.com/microsoft/pyright
[raku]: https://github.com/rakudo/rakudo
[remark-lint]: https://github.com/remarkjs/remark-lint
[renovate-config-validator]: https://github.com/renovatebot/renovate
[revive]: https://github.com/mgechev/revive
[roslynator]: https://github.com/dotnet/Roslynator
[rst-lint]: https://github.com/twolfson/restructuredtext-lint
[rstcheck]: https://github.com/myint/rstcheck
[rubocop]: https://github.com/rubocop-hq/rubocop
[ruff]: https://github.com/astral-sh/ruff
[rustfmt]: https://github.com/rust-lang/rustfmt
[scalafix]: https://github.com/scalacenter/scalafix
[scalafmt]: https://github.com/scalameta/scalafmt
[secretlint]: https://github.com/secretlint/secretlint
[semgrep]: https://github.com/returntocorp/semgrep
[sfdx-scanner-apex]: https://github.com/forcedotcom/sfdx-scanner
[sfdx-scanner-aura]: https://github.com/forcedotcom/sfdx-scanner
[sfdx-scanner-lwc]: https://github.com/forcedotcom/sfdx-scanner
[shellcheck]: https://github.com/koalaman/shellcheck
[shfmt]: https://github.com/mvdan/sh
[snakefmt]: https://github.com/snakemake/snakefmt
[snakemake]: https://github.com/snakemake/snakemake
[spectral]: https://github.com/stoplightio/spectral
[sqlfluff]: https://github.com/sqlfluff/sqlfluff
[standardjs]: https://github.com/standard/standard
[stylelint]: https://github.com/stylelint/stylelint
[swiftlint]: https://github.com/realm/SwiftLint
[syft]: https://github.com/anchore/syft
[tekton-lint]: https://github.com/IBM/tekton-lint
[terraform-fmt]: https://github.com/hashicorp/terraform
[terragrunt]: https://github.com/gruntwork-io/terragrunt
[terrascan]: https://github.com/tenable/terrascan
[textlint]: https://github.com/textlint/textlint
[tflint]: https://github.com/terraform-linters/tflint
[trivy]: https://github.com/aquasecurity/trivy
[trufflehog]: https://github.com/trufflesecurity/trufflehog
[tsqllint]: https://github.com/tsqllint/tsqllint
[v8r]: https://github.com/chris48s/v8r
[vale]: https://github.com/errata-ai/vale
[yamllint]: https://github.com/adrienverge/yamllint
