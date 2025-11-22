# Contribution

Rules for commit messages and branch names to maintain consistency and readability in our project.

## Branch Naming Conventions

All branch names must follow a consistent pattern to clearly indicate their purpose. Use the following prefixes:

* **`feature/*`:** For new features.  Example: `feature/add-user-authentication`
* **`test/*`:** For testing a hypothesis. Example: `test/will-it-work-or-wont`
* **`fix/*`:** For bug fixes. Example: `fix/resolve-login-issue`
* **`refactor/*`:** For rewrite/restructure your code. Example: `refactor/update-react-19`
* **`build/*`:** For changes in build components like build tool, ci pipeline, dependencies. Example: `build/update-deploy-scripts`
* **`ops/*`:** For changes that affect operational components like infrastructure, deployment, backup, recovery. Example: `ops/backup-utils`
* **`docs/*`:** For branches that affect documentation only. Example: `docs/update-contrubution`

**Invalid Branch Names:**

Avoid vague or ambiguous branch names.  Examples of invalid branch names:

* `my-branch`
* `work-in-progress`
* `bugfix`

## Commit Message Conventions

Commit messages should be clear, concise, and informative.  They should follow the following format:

`<type>(Optional: <scope>): <short description>`

Where:

* **`<type>`:**  One of the following types:
  * `feat`: A new feature for the user.
  * `fix`: A bug fix for the user.
  * `hotfix`: A 🔥hot🔥 fix.
  * `docs`: Changes to the documentation.
  * `style`: Changes that don't affect the functionality (e.g., formatting, whitespace).
  * `refactor`: A code change that neither fixes a bug nor adds a feature.
  * `test`: Adding missing tests or refactoring existing tests.
  * `build`: Add command in build script.

* **`<scope>`:** (Optional) A short description of the area of the codebase affected by the change.  Example: `auth`, `database`, `ui`.

* **`<short description>`:** A concise summary of the change (50 characters or less).

**Example Commit Messages:**

* `feat(auth): Add user authentication`
* `fix(ui): Resolve button alignment issue`
* `docs(readme): Update installation instructions`
* `style: Fix formatting inconsistencies and eslint fixes`
* `refactor(database): Improve database query performance`
* `test(api): Add unit tests for API endpoints`

## References

[Conventional Commit Messages^](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13)
