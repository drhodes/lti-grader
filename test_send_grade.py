#!/usr/bin/env python3

import argparse
from answer_set import AnswerSet
from lti import OutcomeRequest
import sys
def json_err(msg): json.dumps({"Ok": False, "Error": msg})

def main():
    consumer = "0c676cd363e7fdb18c3a855eec8ab180213af6c201733bc14ab50c4f37c74b29"
    secret = "0f41e00085b3f411ea5ebb4093c37b56f2bce67bd00bf4a69dffbef2c09bf244"
    service_url = "http://localhost:18000/courses/course-v1:remox+c1+t1/xblock/block-v1:remox+c1+t1+type@lti_consumer+block@13c61a21043e435698ec1303f884a69a/handler_noauth/outcome_service_handler"
    sourcedid = "course-v1%3Aremox%2Bc1%2Bt1:localhost%3A18000-13c61a21043e435698ec1303f884a69a:55fc413721b3acac97cdd98c831b70c4"
    grade = ".4242424242"

    
    outcome_request = OutcomeRequest({
        'consumer_key': consumer,
        'consumer_secret': secret,
        'lis_outcome_service_url': service_url,
        'lis_result_sourcedid': sourcedid,
    })
    
    outcome_response = outcome_request.post_replace_result(grade)
    print(outcome_response)

if __name__ == "__main__":
    print(main())
