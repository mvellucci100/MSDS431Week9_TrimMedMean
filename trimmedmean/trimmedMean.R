# Function to compute trimmed mean in R with 5% trimming from both ends
trimmed_mean <- function(data, trim_percentage = 0.05) {
  # Calculate the trimmed mean with the trim argument
  result <- mean(data, trim = trim_percentage)
  return(result)
}

# Example integer data
int_sample <- sample(1:1000, 100, replace = TRUE)  # Generate 100 random integers

# Example float data
float_sample <- runif(100, 1, 1000)  # Generate 100 random floats

# Compute trimmed means with 5% trimming
trimmed_mean_int <- trimmed_mean(int_sample, 0.05)
trimmed_mean_float <- trimmed_mean(float_sample, 0.05)

# Print the results
cat("Trimmed Mean for Integer Sample (5% trimming): ", trimmed_mean_int, "\n")
cat("Trimmed Mean for Float Sample (5% trimming): ", trimmed_mean_float, "\n")
