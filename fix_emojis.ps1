Get-ChildItem -Path "frontend/src" -Recurse -Include "*.tsx","*.ts" | ForEach-Object {
    $content = Get-Content $_.FullName -Raw
    $content = $content -replace "🔍", ""
    $content = $content -replace "📊", ""
    $content = $content -replace "⏭️", ""
    $content = $content -replace "👤", ""
    $content = $content -replace "📡", ""
    $content = $content -replace "🎯", ""
    $content = $content -replace "⏳", ""
    $content = $content -replace "🔄", ""
    $content = $content -replace "📋", ""
    $content = $content -replace "📝", ""
    $content = $content -replace "🚀", ""
    $content = $content -replace "⚠️", "Warning:"
    $content = $content -replace "❌", ""
    $content = $content -replace "✅", ""
    $content = $content -replace "🔗", ""
    $content = $content -replace "🔵", ""
    Set-Content $_.FullName $content
}
