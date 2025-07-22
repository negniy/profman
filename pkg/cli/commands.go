package cli

import (
	"context"
	"profman/internal/profile"

	"github.com/urfave/cli/v3"
)

func RegCommands() *cli.Command {
	var name, user, project string
	root := &cli.Command{
		Name:        "Profman",
		Usage:       "Утилита для управления YAML-профилями",
		Description: "Профили хранятся в текущей директории в виде файлов <name>.yaml с полями user и project.",
		UsageText:   "profman profile command [command options]",
		Commands: []*cli.Command{
			&cli.Command{
				Name:        "profile",
				Usage:       "Команды для работы с профилями",
				Description: "Создание, получение, перечисление и удаление профилей.",
				UsageText:   "profman profile <create|get|list|delete> [options]",
				Commands: []*cli.Command{
					&cli.Command{
						Name:        "create",
						Usage:       "Создать профиль: --name --user --project",
						Description: "Создаёт файл <name>.yaml и записывает в него поля user и project.",
						UsageText:   "profman profile create --name=<name> --user=<user> --project=<project>",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Required:    true,
								Destination: &name,
								Usage:       "Имя профиля (без расширения .yaml)",
							},
							&cli.StringFlag{
								Name:        "user",
								Required:    true,
								Destination: &user,
								Usage:       "Значение поля user",
							},
							&cli.StringFlag{
								Name:        "project",
								Required:    true,
								Destination: &project,
								Usage:       "Значение поля project",
							},
						},
						Action: func(ctx context.Context, c *cli.Command) error {
							err := profile.CreateProfile(name, user, project)
							if err != nil {
								cli.ErrWriter.Write([]byte(err.Error()))
							}
							return nil
						},
					},
					&cli.Command{
						Name:        "get",
						Usage:       "Показать профиль: --name",
						Description: "Читает файл <name>.yaml и выводит его содержимое.",
						UsageText:   "profman profile get --name=<name>",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Required:    true,
								Destination: &name,
								Usage:       "Имя профиля (без расширения .yaml)",
							},
						},
						Action: func(ctx context.Context, c *cli.Command) error {
							err := profile.GetProfile(name)
							if err != nil {
								cli.ErrWriter.Write([]byte(err.Error()))
							}
							return nil
						},
					},
					&cli.Command{
						Name:        "list",
						Usage:       "Перечислить все профили",
						Description: "Находит все файлы *.yaml в текущей директории и выводит их содержимое.",
						UsageText:   "profman profile list",
						Action: func(ctx context.Context, c *cli.Command) error {
							err := profile.ListProfile()
							if err != nil {
								cli.ErrWriter.Write([]byte(err.Error()))
							}
							return nil
						},
					},
					&cli.Command{
						Name:        "delete",
						Usage:       "Удалить профиль: --name",
						Description: "Удаляет файл <name>.yaml из текущей директории.",
						UsageText:   "profman profile delete --name=<name>",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Required:    true,
								Usage:       "Имя профиля (без расширения .yaml)",
								Destination: &name,
							},
						},
						Action: func(ctx context.Context, c *cli.Command) error {
							err := profile.DeleteProfile(name)
							if err != nil {
								cli.ErrWriter.Write([]byte(err.Error()))
							}
							return nil
						},
					},
				},
			},
		},
	}
	return root
}
