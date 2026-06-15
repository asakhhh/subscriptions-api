#!/usr/bin/env bash
set -euo pipefail

BASE_URL="${BASE_URL:-http://localhost:8080}"
USER_ID="${USER_ID:-60601fee-2bf1-4721-ae6f-7636e79a0cba}"
USER_ID_2="${USER_ID_2:-11111111-1111-1111-1111-111111111111}"

wait_for_api() {
  echo "Waiting for API at ${BASE_URL}..."
  for _ in $(seq 1 60); do
    code="$(curl -s -o /dev/null -w "%{http_code}" "${BASE_URL}/subscriptions" || true)"
    if [[ "${code}" == "400" ]]; then
      echo "API is ready."
      return 0
    fi
    sleep 1
  done
  echo "API did not become ready in time." >&2
  exit 1
}

create_sub() {
  local name="$1"
  local price="$2"
  local user_id="$3"
  local start_date="$4"
  local end_date="${5:-}"

  local payload
  if [[ -n "${end_date}" ]]; then
    payload=$(printf '{"service_name":"%s","price":%s,"user_id":"%s","start_date":"%s","end_date":"%s"}' \
      "${name}" "${price}" "${user_id}" "${start_date}" "${end_date}")
  else
    payload=$(printf '{"service_name":"%s","price":%s,"user_id":"%s","start_date":"%s"}' \
      "${name}" "${price}" "${user_id}" "${start_date}")
  fi

  echo "Creating ${name}..."
  curl -sf -X POST "${BASE_URL}/create_subscription" \
    -H "Content-Type: application/json" \
    -d "${payload}"
  echo
}

wait_for_api

create_sub "Yandex Plus" 400 "${USER_ID}" "07-2025"
create_sub "Netflix" 500 "${USER_ID}" "01-2025" "06-2025"
create_sub "Spotify" 300 "${USER_ID_2}" "03-2025" "12-2025"

echo "Seed complete."
