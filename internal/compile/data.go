package compile

import (
	"html/template"
	"net/url"

	"github.com/NikitaKissa/cvgen/internal/cv"
	tmpl "github.com/NikitaKissa/cvgen/internal/template"
)

// ModelToTemplateData converts the domain model plus a resolved style
// string into the presentation-facing structure that user templates see.
//
// This is the boundary between "what the CV is" (cv.CV) and "what the
// template gets" (tmpl.TemplateData) — it belongs here, in compile, and
// not in internal/template, because internal/template must stay
// unaware of the CV domain model.
func ModelToTemplateData(model cv.CV, style string) tmpl.TemplateData {
	return tmpl.TemplateData{
		Style: template.CSS(style),
		Basics: tmpl.Basics{
			FullName:    model.Basics.FullName,
			Position:    model.Basics.Position,
			Description: model.Basics.Description,
			Email:       model.Basics.Email,
			Phone:       model.Basics.Phone,
			Photo:       template.URL(derefOrEmpty(model.Basics.Photo)),
			Links:       modelLinksToTmplLinks(model.Basics.Links),
		},
		Education:    modelEducationToTmplEducation(model.Education),
		Experience:   modelExperienceToTmplExperience(model.Experience),
		SkillGroups:  model.SkillGroups,
		Certificates: modelCertificatesToTmplCertificates(model.Certificates),
		Projects:     modelProjectsToTmplProjects(model.Projects),
		Languages:    model.Languages,
		Clause:       model.Clause,
	}
}

func modelLinksToTmplLinks(links cv.Links) tmpl.Links {
	tmplLinks := make(tmpl.Links, len(links))
	for name, value := range links {
		tmplLinks[name] = template.URL(value.String())
	}
	return tmplLinks
}

func derefOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefURLToTemplateURL(u *url.URL) *template.URL {
	if u == nil {
		return nil
	}

	tmplURL := template.URL(u.String())
	return &tmplURL
}

func modelEducationToTmplEducation(model []cv.Education) []tmpl.Education {
	result := make([]tmpl.Education, len(model))

	for i, edu := range model {
		result[i] = tmpl.Education{
			Institution:        edu.Institution,
			InstitutionWebsite: derefURLToTemplateURL(edu.InstitutionWebsite),
			FieldOfStudy:       edu.FieldOfStudy,
			Description:        edu.Description,
			From:               edu.From,
			To:                 edu.To,
		}
	}

	return result
}

func modelExperienceToTmplExperience(model []cv.Experience) []tmpl.Experience {
	result := make([]tmpl.Experience, len(model))

	for i, exp := range model {
		result[i] = tmpl.Experience{
			Company:        exp.Company,
			CompanyWebsite: derefURLToTemplateURL(exp.CompanyWebsite),
			Position:       exp.Position,
			Description:    exp.Description,
			Stack:          exp.Stack,
			From:           exp.From,
			To:             exp.To,
		}
	}

	return result
}

func modelCertificatesToTmplCertificates(model []cv.Certificate) []tmpl.Certificate {
	result := make([]tmpl.Certificate, len(model))

	for i, cert := range model {
		result[i] = tmpl.Certificate{
			Name:   cert.Name,
			Issuer: cert.Issuer,
			Url:    template.URL(cert.Url.String()),
		}
	}

	return result
}

func modelProjectsToTmplProjects(model []cv.Project) []tmpl.Project {
	result := make([]tmpl.Project, len(model))

	for i, proj := range model {
		result[i] = tmpl.Project{
			Name:        proj.Name,
			Description: proj.Description,
			Url:         template.URL(proj.Url.String()),
		}
	}

	return result
}
