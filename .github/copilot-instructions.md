# Правила проекта

## Общие правила
- Всегда отвечай по-русски.
- Документация, комментарии, PR и коммиты — максимально кратко и по-человечески.
- Не дублируй README и `docs/`; если есть подробное описание, ссылайся на него.
- Перед правками смотри соседний код и сохраняй стиль файла. Лишний рефакторинг не делай.

## Архитектура
- `frontend/` — Nuxt 4 dashboard. Клиентский код лежит в `frontend/app/`, локальные API-роуты — в `frontend/server/api/`.
- `backend/` — Go API на `net/http`. Роутинг и handlers находятся в `backend/internal/http/`, точка входа — `backend/cmd/api/main.go`.
- `worker/` — отдельный Go worker, точка входа — `worker/cmd/worker/main.go`.
- `infra/` — Docker, compose и env-конфиги.

## Сборка и проверка
- Frontend: `pnpm install`, `pnpm dev`, `pnpm build`, `pnpm preview`, `pnpm lint`, `pnpm typecheck`.
- Backend: `cd backend && go build ./cmd/api`, `cd backend && go test ./...`.
- Worker: `cd worker && go build ./cmd/worker`.
- В frontend используй `pnpm@10.33.0`; в Go-модулях версия `1.26.1`.

## Конвенции
- Ветки: `feature/...`, `fix/...`, `docs/...`, `chore/...`.
- В `main` напрямую не пушим.
- PR — через squash merge; комментарии перед merge закрываем.
- В backend без необходимости не добавляй фреймворки и внешние зависимости.
- В frontend держись существующих паттернов Nuxt, composables, pages/components и темы в `frontend/app/assets/css/main.css`.