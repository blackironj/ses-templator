package ses

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	sessdk "github.com/aws/aws-sdk-go/service/ses"
)

func GetSESTemplateList(page int) ([]*sessdk.TemplateMetadata, error) {
	if EamilServiceSess == nil {
		return nil, errors.New("fail to access")
	}
	sesClient := sessdk.New(EamilServiceSess)

	itemsFrom := (page * 10) - 9

	listTemplatesInput := sessdk.ListTemplatesInput{
		MaxItems: aws.Int64(int64(itemsFrom)),
	}

	listTemplatesOutput, err := sesClient.ListTemplates(&listTemplatesInput)
	if err != nil {
		return nil, err
	}
	return listTemplatesOutput.TemplatesMetadata, nil
}

func CreateSESTemplate(name, subj, htmlbody *string) error {
	sesClient := sessdk.New(EamilServiceSess)

	templ := &sessdk.Template{
		TemplateName: name,
		SubjectPart:  subj,
		TextPart:     htmlbody,
	}

	createTemplateInput := &ses.CreateTemplateInput{
		Template: templ,
	}

	_, err := sesClient.CreateTemplate(createTemplateInput)
	if err != nil {
		return err
	}
	return nil
}
