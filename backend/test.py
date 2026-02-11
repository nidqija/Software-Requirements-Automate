import requests
import json

if __name__ == "__main__":
   
    number_of_questions = 10
    question_number = 0

    while question_number < number_of_questions:
        question = str(input("Enter your question: "))
        data = {
            "model": "gemma3:1b",
            "messages": [
                {
                    "role": "user",
                    "content": question
                }
            ],
            "stream": False
        }
        url = "http://localhost:11434/api/chat"
        response = requests.post(url, json=data)
        response_json = json.loads(response.text)
        ai_reply = response_json['message']['content']
        print(ai_reply)
        question_number += 1
        


