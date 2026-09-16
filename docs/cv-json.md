# CV JSON Format

`cvgen` uses JSON to describe the content of a CV. The JSON is independent of the template and controls the data that is available to the renderer.

## Basic structure

A CV is represented by the following top-level sections:

```json
{
  "basics": {},
  "education": [],
  "experience": [],
  "skill_groups": [],
  "certificates": [],
  "projects": [],
  "languages": [],
  "clause": null
}
```

All field names use `snake_case`.

## Nullable fields

Fields corresponding to nullable values may either contain a value or `null`.

For optional scalar fields, omitting the field and setting it to `null` both represent an absent value.

For example:

```json
{
  "email": null,
  "phone": null
}
```

Time periods follow the same principle. The `to` field may either be omitted or set to `null`. Both mean that the period is ongoing (`Present`).

## Dates

Dates use the `YYYY-MM` format:

```json
{
  "from": "2025-08",
  "to": "2026-08"
}
```

For an ongoing period, `to` can be omitted or set to `null`:

```json
{
  "from": "2025-08"
}
```

or:

```json
{
  "from": "2025-08",
  "to": null
}
```

Both represent a period from August 2025 to the present.

## Basics

The `basics` section contains the main personal and contact information.

| Field         | Type             | Description                             |
| ------------- | ---------------- | --------------------------------------- |
| `photo`       | `string \| null` | Path or URL of the profile image.       |
| `full_name`   | `string`         | Full name.                              |
| `position`    | `string`         | Professional title or desired position. |
| `description` | `string`         | Short professional description.         |
| `email`       | `string \| null` | Email address.                          |
| `phone`       | `string \| null` | Phone number.                           |
| `links`       | `object`         | Named external links.                   |

### Links

`links` is an object where each key is a user-defined label and its value is a URL.

```json
{
  "links": {
    "LinkedIn": "https://www.linkedin.com/in/example",
    "GitHub": "https://github.com/example"
  }
}
```

The labels are not restricted to a predefined set.

## Education

`education` is an array of educational experiences.

| Field                 | Type             | Description                                              |
| --------------------- | ---------------- | -------------------------------------------------------- |
| `institution`         | `string`         | Name of the educational institution.                     |
| `institution_website` | `string \| null` | Institution website.                                     |
| `field_of_study`      | `string`         | Field of study.                                          |
| `description`         | `string \| null` | Additional information about the education.              |
| `from`                | `string`         | Start date (`YYYY-MM`).                                  |
| `to`                  | `string \| null` | End date (`YYYY-MM`). Omitted or `null` means `Present`. |

Example:

```json
{
  "institution": "Elektroniczne Zakady Naukowe",
  "institution_website": "https://ezn.edu.pl",
  "field_of_study": "Computer science",
  "description": "The best tech-school in the state",
  "from": "2023-09",
  "to": "2028-06"
}
```

## Experience

`experience` is an array of professional experiences.

| Field             | Type             | Description                                              |
| ----------------- | ---------------- | -------------------------------------------------------- |
| `company`         | `string`         | Company or organization name.                            |
| `company_website` | `string \| null` | Company website.                                         |
| `position`        | `string`         | Position or job title.                                   |
| `description`     | `string \| null` | Description of the work.                                 |
| `stack`           | `string[]`       | Technologies, tools, or skills used.                     |
| `from`            | `string`         | Start date (`YYYY-MM`).                                  |
| `to`              | `string \| null` | End date (`YYYY-MM`). Omitted or `null` means `Present`. |

Example:

```json
{
  "company": "Letiko",
  "company_website": "https://letiko.com",
  "position": "Frontend developer",
  "description": "Developing of CRM page for CarAssistant",
  "stack": ["Nuxt", "TypeScript", "Tailwind CSS"],
  "from": "2025-08",
  "to": "2026-08"
}
```

## Skill groups

`skill_groups` contains named groups of skills.

| Field    | Type       | Description                    |
| -------- | ---------- | ------------------------------ |
| `name`   | `string`   | Name of the skill group.       |
| `skills` | `string[]` | Skills belonging to the group. |

Example:

```json
{
  "name": "Hard skills",
  "skills": [
    "Git",
    "Golang development",
    "PostgreSQL"
  ]
}
```

Group names are user-defined and are not restricted to values such as `Hard skills` or `Soft skills`.

## Certificates

`certificates` contains professional or educational certificates.

| Field    | Type     | Description                                          |
| -------- | -------- | ---------------------------------------------------- |
| `name`   | `string` | Certificate name.                                    |
| `issuer` | `string` | Organization that issued the certificate.            |
| `url`    | `string` | URL where the certificate can be viewed or verified. |

Example:

```json
{
  "name": "CCNA: Introduction to Networks",
  "issuer": "Cisco",
  "url": "https://example.com/certificate"
}
```

## Projects

`projects` contains personal, professional, or open-source projects.

| Field         | Type             | Description          |
| ------------- | ---------------- | -------------------- |
| `name`        | `string`         | Project name.        |
| `description` | `string \| null` | Project description. |
| `url`         | `string`         | Project URL.         |

Example:

```json
{
  "name": "cvgen",
  "description": "Generator CV in HTML format from JSON",
  "url": "https://github.com/NikitaKissa/cvgen"
}
```

## Languages

`languages` contains spoken or written languages and their proficiency levels.

| Field         | Type     | Description                       |
| ------------- | -------- | --------------------------------- |
| `name`        | `string` | Language name.                    |
| `proficiency` | `string` | Proficiency level or description. |

`proficiency` is intentionally represented as a string, so the value is not restricted to a specific grading system.

Example:

```json
{
  "name": "English",
  "proficiency": "B2"
}
```

Other values such as `Native`, `Fluent`, or custom descriptions are also valid.

## Clause

`clause` is an optional text block, typically used for a legal or recruitment consent clause.

It can contain any text or be `null` / omitted.

Example:

```json
{
  "clause": "I consent to the processing of my personal data..."
}
```

## Images

The `basics.photo` field can contain either a local file path or an external URL.

For a local image, `cvgen` embeds the image into the generated HTML as a Base64 data URI. This keeps the generated document self-contained.

For an image specified by URL, the generated HTML keeps the external URL. This produces a smaller HTML document, but the image must remain available at that URL.

Example:

```json
{
  "basics": {
    "photo": "/home/user/photos/profile.jpg"
  }
}
```

or:

```json
{
  "basics": {
    "photo": "https://example.com/profile.jpg"
  }
}
```

## Complete example

A complete input file can look like this:

```json
{
  "basics": {
    "photo": "/home/user/photo.jpg",
    "full_name": "Mykyta Kissa",
    "position": "Backend Go Developer",
    "description": "Highly motivated junior specialist",
    "email": "test@test.com",
    "phone": "+48123123123",
    "links": {
      "LinkedIn": "https://www.linkedin.com/in/example",
      "Credly": "https://www.credly.com/users/example"
    }
  },
  "education": [
    {
      "institution": "Elektroniczne Zakady Naukowe",
      "institution_website": "https://ezn.edu.pl",
      "field_of_study": "Computer Science",
      "description": "The best tech-school in the state",
      "from": "2023-09",
      "to": "2028-06"
    }
  ],
  "experience": [
    {
      "company": "Letiko",
      "company_website": "https://letiko.com",
      "position": "Frontend Developer",
      "description": "Developing a CRM page for CarAssistant",
      "stack": [
        "Nuxt",
        "TypeScript",
        "Tailwind CSS"
      ],
      "from": "2025-08",
      "to": "2026-08"
    }
  ],
  "skill_groups": [
    {
      "name": "Hard skills",
      "skills": [
        "Git",
        "Golang development",
        "PostgreSQL",
        "HTML, CSS, JS"
      ]
    },
    {
      "name": "Soft skills",
      "skills": [
        "Teamwork",
        "Time management"
      ]
    }
  ],
  "certificates": [
    {
      "name": "CCNA: Switching, Routing, and Wireless Essentials",
      "issuer": "Cisco",
      "url": "https://example.com/certificate"
    }
  ],
  "projects": [
    {
      "name": "cvgen",
      "description": "Generator CV in HTML format from JSON",
      "url": "https://github.com/NikitaKissa/cvgen"
    },
    {
      "name": "CarAssistant CRM",
      "url": "https://carassistant.ai"
    }
  ],
  "languages": [
    {
      "name": "Ukrainian",
      "proficiency": "Native"
    },
    {
      "name": "English",
      "proficiency": "B2"
    }
  ],
  "clause": "I consent to the processing of my personal data for recruitment purposes."
}
```
