import json

# include the answer in the problem, but symmetric encrypt it first
# with $(SECRET_KEY)

def check_function(e, ans):
    """
    'response' is a dictionary that contains two keys, 'answer' and
    'state'.

    The value of "answer" is the JSON string that "getGrade" returns.
    The value of "state" is the JSON string that "getState" returns.
    Clicking either "Submit" or "Save" registers the current state.
    """
    response = json.loads(ans)

    # You can use the value of the answer key to grade:
    answer = json.loads(response["answer"])
    return answer == "correct"

    # Or you can use the value of the state key to grade:
    """
    state = json.loads(response["state"])
    return state["selectedChoice"] == "correct"
    """
