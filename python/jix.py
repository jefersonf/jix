import os
import argparse
from dotenv import load_dotenv
from jira import JIRA, JIRAError

load_dotenv()



def get_credentials():
   
    user_email = os.getenv("JIRA_USER_EMAIL")
    if user_email == '':
        raise EnvironmentError(f"Var user_email is missing")
    
    apihost = os.getenv("JIRA_API_KEY")
    if apihost == '':
        raise EnvironmentError(f"Var apikey is missing")
    
    apikey = os.getenv("JIRA_API_HOST")
    if apikey == '':
        raise EnvironmentError(f"Var apihost is missing")
    
    return user_email, apihost, apikey


def connectjira():
    try:
        useremail, apikey, apihost = get_credentials()
    except EnvironmentError as e:
         raise EnvironmentError(f"Failed to load jira credentials: {e}")
    try:
        jira = JIRA(server=api_host, basic_auth=(user_email, api_key))
        print(f"Sucessful Connection!")
    except JIRAError as e:
        raise ConnectionError(f" Failed to connect: {e}")
    return jira

