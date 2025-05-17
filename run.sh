#!/usr/bin/env bash

set -e

export OLLAMA_HOST="localhost"
# OLLAMA_PORT="11434"
# OLLAMA_API="http://$OLLAMA_HOST:$OLLAMA_PORT/v1/models"
# OLLAMA_MODEL="phi4"

# Step 1: Generate prompts for all categories and topics
echo "🔧 Generating prompts for all combinations..."
categories=("classification" "demo" "topics")
for category in "${categories[@]}"; do
  for topic_dir in "resources/$category"/*; do
    if [ -d "$topic_dir" ]; then
      topic=$(basename "$topic_dir")
      echo "🔧 Generating prompt for category=$category, topic=$topic"
      if [ "$category" = "classification" ]; then
        .build/ai-explorer prompt \
          --category="$category" \
          --topic="$topic" \
          --query="Explain Sliding Window Protocol to a CCIE person, keep it short"
      else
        .build/ai-explorer prompt \
          --category="$category" \
          --topic="$topic"
      fi
    fi
  done
done

# Step 2: Generate the llm
echo "🔧 Generating llm..."
.build/ai-explorer llm  \
    --model=phi4 \
    --provider=ollama \
    --prompt="resources/classification/router/prompt.txt" \
    --temperature=0.8 \
    --timeout=2m \
    --output="resources/classification/router/llm.txt"
