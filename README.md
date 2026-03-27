# DropIt

DropIt — простая deploy-платформа, вдохновленная Vercel, для небольших сервисов и Telegram-ботов на одном VPS.

## Stack
- Frontend: Vue 3 + TypeScript
- Backend: Go
- Worker: Go
- DB: PostgreSQL
- Containers: Docker
- Proxy: Nginx

## MVP
- auth
- projects
- deploy via Dockerfile
- logs
- env variables
- deployment history
- restart / stop
- subdomain routing
- healthcheck

## Repository structure
- frontend/ — dashboard
- backend/ — API
- worker/ — deploy/build worker
- infra/ — infrastructure config
- docs/ — project docs

## Workflow

### Основные правила
- В ветку `main` напрямую **не пушим**
- Вся работа ведётся через отдельные ветки
- Все изменения проходят через **Pull Request (PR)**
- Минимум **1 одобрение (review)** перед merge

---

### Именование веток
Используем понятные префиксы:

- `feature/...` — новая функциональность  
- `fix/...` — исправление багов  
- `docs/...` — изменения в документации  
- `chore/...` — технические задачи (рефакторинг, настройки и т.д.)

**Примеры:**
- `feature/auth-api`
- `feature/project-dashboard`
- `fix/deploy-error`
- `docs/readme-update`

---

### Процесс работы
1. Создаёте ветку от `main`
2. Делаете изменения
3. Пушите ветку в репозиторий
4. Открываете Pull Request в `main`
5. Кто-то из команды делает review
6. После одобрения — merge

---

### Merge
- Предпочтительно использовать **Squash merge**
- Один PR = один коммит в `main`

---

### Важно
- Если в PR есть комментарии — их нужно закрыть перед merge
- Не делаем force push в `main`
- Стараемся делать небольшие и понятные PR
