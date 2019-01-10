output "repository_url" {
  description = "Git Repository URL"
  value       = "${google_sourcerepo_repository.new_git_repository.url}"
}