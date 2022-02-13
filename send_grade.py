#!/usr/bin/env python3

import argparse
from answer_set import AnswerSet
from lti import OutcomeRequest
import sys

parser = argparse.ArgumentParser(description='grade_render')
parser.add_argument('--consumer', dest='consumer', default="", help="", required=True)
parser.add_argument('--secret', dest='secret', default="", help="", required=True)
parser.add_argument('--lis-outcome-service-url', dest='lis_outcome_service_url', default="", help="", required=True)
parser.add_argument('--lis-result-sourcedid', dest='lis_result_sourcedid', default="", help="", required=True)
parser.add_argument('--grade', dest='grade', default="", help="", required=True)

args = parser.parse_args()

def json_err(msg): json.dumps({"Ok": False, "Error": msg})

def main():
    # TODO can this constructor throw an exceptions?
    outcome_request = OutcomeRequest({
        'consumer_key': args.consumer,
        'consumer_secret': args.secret,
        'lis_outcome_service_url': args.lis_outcome_service_url,
        'lis_result_sourcedid': args.lis_result_sourcedid,
    })
    
    outcome_response = outcome_request.post_replace_result(args.grade)
    print(outcome_response)

if __name__ == "__main__":
    print(main())
