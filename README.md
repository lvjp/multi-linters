# 🐐 Multiple linter on multiple architectures

## Comparison

| Linter                                                       |  multi-linter (2)  | mega-linter (110)  | super-linter (75)  |
| ------------------------------------------------------------ | :----------------: | :----------------: | :----------------: |
| [actionlint][actionlint]                                     |                    | :white_check_mark: | :white_check_mark: |
| [ansible-lint][ansible-lint]                                 |                    | :white_check_mark: | :white_check_mark: |
| [Azure Resource Manager Template Toolkit (arm-ttk)][arm-ttk] |                    | :white_check_mark: | :white_check_mark: |
| [ASL Validator][asl-validator]                               |                    |                    | :white_check_mark: |
| [bandit][bandit]                                             |                    | :white_check_mark: |                    |
| bash-exec                                                    |                    | :white_check_mark: | :white_check_mark: |
| [bicep-linter][bicep-linter]                                 |                    | :white_check_mark: |                    |
| [black][black]                                               |                    | :white_check_mark: | :white_check_mark: |
| [AWS CloudFormation Linter (cfn-lint)][cfn-lint]             |                    | :white_check_mark: | :white_check_mark: |
| [checkmake][checkmake]                                       |                    | :white_check_mark: |                    |
| [checkov][checkov]                                           |                    | :white_check_mark: | :white_check_mark: |
| [checkstyle][checkstyle]                                     |                    | :white_check_mark: | :white_check_mark: |
| [chktex][chktex]                                             |                    |                    | :white_check_mark: |
| [clang-format][clang-format]                                 |                    | :white_check_mark: | :white_check_mark: |
| [clippy][clippy]                                             |                    | :white_check_mark: | :white_check_mark: |
| [clj-kondo][clj-kondo]                                       |                    | :white_check_mark: | :white_check_mark: |
| [cljstyle][cljstyle]                                         |                    | :white_check_mark: |                    |
| [coffeelint][coffeelint]                                     |                    | :white_check_mark: | :white_check_mark: |
| [commitlint][commitlint]                                     |                    |                    | :white_check_mark: |
| [cpplint][cpplint]                                           |                    | :white_check_mark: | :white_check_mark: |
| [csharpier][csharpier]                                       |                    | :white_check_mark: |                    |
| [cspell][cspell]                                             |                    | :white_check_mark: |                    |
| [dart-analyze][dart-analyze]                                 |                    | :white_check_mark: | :white_check_mark: |
| [detekt][detekt]                                             |                    | :white_check_mark: |                    |
| [devskim][devskim]                                           |                    | :white_check_mark: |                    |
| [djlint][djlint]                                             |                    | :white_check_mark: |                    |
| [dotenv-linter][dotenv-linter]                               |                    | :white_check_mark: | :white_check_mark: |
| [dotnet-format][dotnet-format]                               |                    | :white_check_mark: | :white_check_mark: |
| [dustilock][dustilock]                                       |                    | :white_check_mark: |                    |
| [editorconfig-checker][editorconfig-checker]                 |                    | :white_check_mark: | :white_check_mark: |
| [eslint-plugin-json][eslint-plugin-json]                     |                    |                    | :white_check_mark: |
| [eslint-plugin-jsonc][eslint-plugin-jsonc]                   |                    | :white_check_mark: | :white_check_mark: |
| [eslint-plugin-jsx-a11y][eslint-plugin-jsx-a11y]             |                    | :white_check_mark: | :white_check_mark: |
| [eslint-plugin-react][eslint-plugin-react]                   |                    | :white_check_mark: | :white_check_mark: |
| [eslint-typescript][eslint-typescript]                       |                    | :white_check_mark: | :white_check_mark: |
| [ESLint][eslint]                                             |                    | :white_check_mark: | :white_check_mark: |
| [flake8][flake8]                                             |                    | :white_check_mark: | :white_check_mark: |
| [gherkin-lint][gherkin-lint]                                 |                    | :white_check_mark: | :white_check_mark: |
| [Git conflict markers presence in files][git-diff]           |                    | :white_check_mark: | :white_check_mark: |
| [GitLeaks][gitleaks]                                         |                    | :white_check_mark: | :white_check_mark: |
| [golangci-lint][golangci-lint]                               | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| [google-java-format][google-java-format]                     |                    |                    | :white_check_mark: |
| [GoReleaser][goreleaser]                                     |                    |                    | :white_check_mark: |
| [graphql-schema-linter][graphql-schema-linter]               |                    | :white_check_mark: |                    |
| [grype][grype]                                               |                    | :white_check_mark: |                    |
| [Haskell Dockerfile Linter][hadolint]                        | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| [helm][helm]                                                 |                    | :white_check_mark: |                    |
| [HTMLHint][htmlhint]                                         |                    | :white_check_mark: | :white_check_mark: |
| [isort][isort]                                               |                    | :white_check_mark: | :white_check_mark: |
| [jscpd][jscpd]                                               |                    | :white_check_mark: | :white_check_mark: |
| [jsonlint][jsonlint]                                         |                    | :white_check_mark: |                    |
| [kics][kics]                                                 |                    | :white_check_mark: |                    |
| [ktlint][ktlint]                                             |                    | :white_check_mark: | :white_check_mark: |
| [kubeconform][kubeconform]                                   |                    | :white_check_mark: | :white_check_mark: |
| [kubescape][kubescape]                                       |                    | :white_check_mark: |                    |
| [LibXML][libxml]                                             |                    |                    | :white_check_mark: |
| [lightning-flow-scanner][lightning-flow-scanner]             |                    | :white_check_mark: |                    |
| [lintr][lintr]                                               |                    | :white_check_mark: | :white_check_mark: |
| [ls-lint][ls-lint]                                           |                    | :white_check_mark: |                    |
| [luacheck][luacheck]                                         |                    | :white_check_mark: | :white_check_mark: |
| [lychee][lychee]                                             |                    | :white_check_mark: |                    |
| [markdown-link-check][markdown-link-check]                   |                    | :white_check_mark: |                    |
| [markdown-table-formatter][markdown-table-formatter]         |                    | :white_check_mark: |                    |
| [markdownlint][markdownlint]                                 |                    | :white_check_mark: | :white_check_mark: |
| [mypy][mypy]                                                 |                    | :white_check_mark: |                    |
| [npm-groovy-lint][npm-groovy-lint]                           |                    | :white_check_mark: | :white_check_mark: |
| [npm-package-json-lint][npm-package-json-lint]               |                    | :white_check_mark: |                    |
| [perlcritic][perlcritic]                                     |                    | :white_check_mark: | :white_check_mark: |
| [php-cs-fixer][php-cs-fixer]                                 |                    | :white_check_mark: |                    |
| [PHP CodeSniffer][phpcs]                                     |                    | :white_check_mark: | :white_check_mark: |
| [phplint][phplint]                                           |                    | :white_check_mark: |                    |
| [PHPStan][phpstan]                                           |                    | :white_check_mark: | :white_check_mark: |
| [PHP built-in linter][phpstd]                                |                    |                    | :white_check_mark: |
| [pmd][pmd]                                                   |                    | :white_check_mark: |                    |
| [powershell-formatter][powershell-formatter]                 |                    | :white_check_mark: |                    |
| [powershell][powershell]                                     |                    | :white_check_mark: |                    |
| [Prettier][prettier]                                         |                    | :white_check_mark: | :white_check_mark: |
| [proselint][proselint]                                       |                    | :white_check_mark: |                    |
| [protolint][protolint]                                       |                    | :white_check_mark: | :white_check_mark: |
| [PSScriptAnalyzer][ps-script-analyzer]                       |                    |                    | :white_check_mark: |
| [Psalm][psalm]                                               |                    | :white_check_mark: | :white_check_mark: |
| [puppet-lint][puppet-lint]                                   |                    | :white_check_mark: |                    |
| [pyink][pyink]                                               |                    |                    | :white_check_mark: |
| [pylint][pylint]                                             |                    | :white_check_mark: | :white_check_mark: |
| [pyright][pyright]                                           |                    | :white_check_mark: |                    |
| [Raku][raku]                                                 |                    | :white_check_mark: | :white_check_mark: |
| [remark-lint][remark-lint]                                   |                    | :white_check_mark: |                    |
| [renovate-config-validator][renovate-config-validator]       |                    |                    | :white_check_mark: |
| [revive][revive]                                             |                    | :white_check_mark: |                    |
| [roslynator][roslynator]                                     |                    | :white_check_mark: |                    |
| [rst-lint][rst-lint]                                         |                    | :white_check_mark: |                    |
| [rstcheck][rstcheck]                                         |                    | :white_check_mark: |                    |
| [RuboCop][rubocop]                                           |                    | :white_check_mark: | :white_check_mark: |
| [ruff][ruff]                                                 |                    | :white_check_mark: | :white_check_mark: |
| [Rustfmt][rustfmt]                                           |                    |                    | :white_check_mark: |
| [scalafix][scalafix]                                         |                    | :white_check_mark: |                    |
| [scalafmt][scalafmt]                                         |                    |                    | :white_check_mark: |
| [secretlint][secretlint]                                     |                    | :white_check_mark: |                    |
| [semgrep][semgrep]                                           |                    | :white_check_mark: |                    |
| [sfdx-scanner-apex][sfdx-scanner-apex]                       |                    | :white_check_mark: |                    |
| [sfdx-scanner-aura][sfdx-scanner-aura]                       |                    | :white_check_mark: |                    |
| [sfdx-scanner-lwc][sfdx-scanner-lwc]                         |                    | :white_check_mark: |                    |
| [ShellCheck][shellcheck]                                     |                    | :white_check_mark: | :white_check_mark: |
| [shfmt][shfmt]                                               |                    | :white_check_mark: | :white_check_mark: |
| [snakefmt][snakefmt]                                         |                    | :white_check_mark: | :white_check_mark: |
| [snakemake][snakemake]                                       |                    | :white_check_mark: | :white_check_mark: |
| [spectral][spectral]                                         |                    | :white_check_mark: | :white_check_mark: |
| [sqlfluff][sqlfluff]                                         |                    | :white_check_mark: | :white_check_mark: |
| [standardjs][standardjs]                                     |                    | :white_check_mark: | :white_check_mark: |
| [stylelint][stylelint]                                       |                    | :white_check_mark: | :white_check_mark: |
| [swiftlint][swiftlint]                                       |                    | :white_check_mark: |                    |
| [syft][syft]                                                 |                    | :white_check_mark: |                    |
| [tekton-lint][tekton-lint]                                   |                    | :white_check_mark: | :white_check_mark: |
| [terraform-fmt][terraform-fmt]                               |                    | :white_check_mark: | :white_check_mark: |
| [terragrunt][terragrunt]                                     |                    | :white_check_mark: | :white_check_mark: |
| [terrascan][terrascan]                                       |                    | :white_check_mark: | :white_check_mark: |
| [textlint][textlint]                                         |                    |                    | :white_check_mark: |
| [tflint][tflint]                                             |                    | :white_check_mark: | :white_check_mark: |
| [trivy][trivy]                                               |                    | :white_check_mark: |                    |
| [trufflehog][trufflehog]                                     |                    | :white_check_mark: |                    |
| [tsqllint][tsqllint]                                         |                    | :white_check_mark: |                    |
| [v8r][v8r]                                                   |                    | :white_check_mark: |                    |
| [vale][vale]                                                 |                    | :white_check_mark: |                    |
| [YamlLint][yamllint]                                         |                    | :white_check_mark: | :white_check_mark: |

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
[chktex]: https://git.savannah.nongnu.org/cgit/chktex.git
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
