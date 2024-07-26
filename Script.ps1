# Функция для проверки здоровья сервера
function CheckServerHealth {
    $healthResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/health" -Method GET
    $healthStatus = $healthResponse.Content
    Write-Host "Server health check: $healthStatus"
}

# Функция для регистрации пользователя
function RegisterUser {
    $body = @{
        email = "mail1@example.com"
        password = "1password1"
        username = "user1"
        role = "user"
        first_name = "FirstName"
        last_name = "LastName"
        name = "FullName"
    } | ConvertTo-Json

    try {
        $response = Invoke-WebRequest -Uri "http://localhost:8080/api/register" `
            -Method POST `
            -ContentType "application/json" `
            -Body $body

        $content = $response.Content | ConvertFrom-Json
        return $content.token
    } catch {
        if ($_ -match "duplicate key value violates unique constraint") {
            Write-Host "User already exists, attempting to login"
            return $null
        } else {
            Write-Host "Registration failed: $_"
            return $null
        }
    }
}

# Функция для логина и получения токена
function GetAuthToken {
    $body = @{
        email = "mail1@example.com"
        password = "1password1"
    } | ConvertTo-Json

    try {
        $loginResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/login" `
            -Method POST `
            -ContentType "application/json" `
            -Body $body

        $loginContent = $loginResponse.Content | ConvertFrom-Json
        return $loginContent.token
    } catch {
        Write-Host "Login failed: $_"
        return $null
    }
}

# Функция для получения списка пользователей
function GetUsers {
    param (
        [string]$token
    )

    try {
        $usersResponse = Invoke-WebRequest -Uri "http://localhost:8080/users" `
            -Method GET `
            -Headers @{ Authorization = "Bearer $token" }

        $usersContent = $usersResponse.Content | ConvertFrom-Json
        $users = $usersContent | ConvertTo-Json -Depth 100
        Write-Host "Users: $users"
    } catch {
        Write-Host "Failed to get users: $_"
    }
}

# Основной процесс
CheckServerHealth

$token = RegisterUser
if (-not $token) {
    $token = GetAuthToken
}

if ($token) {
    Write-Host "Obtained token: $token"
    GetUsers -token $token
} else {
    Write-Host "Failed to obtain token"
}
