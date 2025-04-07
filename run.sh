#!/usr/bin/env bash

set -e

# Step 1: Generate the prompt
echo "🔧 Generating prompt..."
.build/ai-explorer prompt \
  --category=classification \
  --topic=router \
  --template="resources/classification/router/template.yaml" \
  --config="resources/classification/router/config.yaml" \
  --query="Explain Sliding Window Protocol to a CCIE person, keep it short"

# Step 2: Generate the llm
echo "🔧 Generating llm..."
.build/ai-explorer llm  \
    --model=phi4 \
    --provider=ollama \
    --prompt="resources/classification/router/prompt.txt" \
    --temperature=0.8 \
    --timeout=2m \
    --output="resources/classification/router/llm.txt"
