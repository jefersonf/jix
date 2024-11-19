import os
import json
import argparse
from dotenv import load_dotenv
from jira import JIRA, JIRAError

load_dotenv()



def get_credentials():
   
    user_email = os.getenv("JIRA_USER_EMAIL")
    if user_email == '':
        raise EnvironmentError(f"Var user_email is missing")
    
    apihost = os.getenv("JIRA_API_HOST")
    if apihost == '':
        raise EnvironmentError(f"Var apikey is missing")
    
    apikey = os.getenv("JIRA_API_KEY")
    if apikey == '':
        raise EnvironmentError(f"Var apihost is missing")
    
    return user_email, apihost, apikey


def connect_to_jira():
    try:
        user_email, apihost, apikey = get_credentials()
    except EnvironmentError as e:
         raise EnvironmentError(f"Failed to load jira credentials: {e}")
    try:
        jira = JIRA(server=apihost, basic_auth=(user_email, apikey))
        print(f"Sucessful Connection!")
    except JIRAError as e:
        raise ConnectionError(f" Failed to connect: {e}")
    return jira

def fetch_issues(project_key,jira):
    issues = jira.search_issues(f'project={project_key}',maxResults=1000)
    return issues

def save_to_jsonl(issues,output_path,project_key):
    json_file_path = os.path.join(output_path, f"{project_key.lower()}.jsonl")
    os.makedirs(os.path.dirname(json_file_path), exist_ok=True)  
    with open(json_file_path, "w") as file:
        for issue in issues:
            json_data = json.dumps(issue)  
            file.write(json_data + "\n") 



def save_to_file(issues, output_format):
    if output_format == 'jsonl':
        save_to_jsonl(issues)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Jira Issue eXtractor")
    parser.add_argument("-p", "--project-key", required=True, help="Jira project key")
    parser.add_argument("-f", "--format", default="jsonl", choices=["jsonl", "csv"], help="Output format")
    parser.add_argument("-o", "--output-path", default="./", help="Path to the output file")
    parser.add_argument("-v", "--verbose", action="store_true", help="Set verbose mode")

    args = parser.parse_args()
    jira = connect_to_jira()
    issues = fetch_issues(args.project_key, jira)
