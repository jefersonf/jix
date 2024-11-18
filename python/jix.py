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


