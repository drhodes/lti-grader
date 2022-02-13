#!/usr/bin/env python3

import sys
import argparse
from mako.template import Template
import json

from answer_set import AnswerSet
import util
import traceback

parser = argparse.ArgumentParser(description='grade_render')
parser.add_argument('--student-answers', dest='student_answers', default="{}", help='json student answers')
parser.add_argument('--staff-answers', dest='staff_answers', default="{}", help='json staff answers')

args = parser.parse_args()

def json_err(msg):
    json.dumps({"Ok": False, "Error": msg, "Html": "", "Grade": 0})

def main():
    # TODO handle exceptions these two lines might throw.
    studentSet = AnswerSet(args.student_answers)
    staffSet = AnswerSet(args.staff_answers)

    if staffSet.size() == 0:
        msg = "--staff-answers were not provided to the render.py program on the answer-server"
        return (False, json_err(msg))

    # try:
    #     grade = staffSet.num_shared_values(studentSet) / staffSet.size()
    # except:
    #     msg = traceback.format_exc() + " \nplease report this to the community TA, thanks!" 
    #     return (False, json_err(msg))
    grade = staffSet.num_shared_values(studentSet) / staffSet.size()

    
    try: 
        html = render(staffSet, studentSet)
        body = json.dumps({"ok": True, "grade": grade, "html": html})
        return (True, body)
    except:
        msg = traceback.format_exc() + " \nplease report this to the community TA, thanks!"        
        return (False, json_err(msg))

def render(staffSet, studentSet):
    answers_html = staffSet.render(studentSet)
    css_styles = open("static/css/remoxblock.css").read()

    ctx = {
        "css_styles": css_styles,
        "answers_html": answers_html,
    }
    return Template(open("static/html/one-page.html").read()).render(ctx=ctx)

if __name__ == "__main__":
    ok, result = main()
    print(result)
    
    if not ok: sys.exit(1)
