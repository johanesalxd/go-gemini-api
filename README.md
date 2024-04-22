https://ai.google.dev/gemini-api/docs/get-started/go

curl http://localhost:8080/generatetext \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"prompt_input": "hi, hello!","model": "gemini-pro"}'

curl http://localhost:8080/chat \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"prompt_input": "hi, hello my name is tommy!","model": "gemini-pro"}'

curl http://localhost:8080/chat \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"prompt_input": "who am i?","model": "gemini-pro"}'