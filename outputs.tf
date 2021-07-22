output "bucket_name" {
  value = google_storage_bucket.bucket.name
}

output "bucket_project" {
  value = google_storage_bucket.bucket.project
}

output "bucket_location" {
  value = google_storage_bucket.bucket.location
}
