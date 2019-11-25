# -*- coding: utf-8 -*-

from eshelper import ESHelper, UtilitiesTracking

def lambda_handler(event, context):

    esh = ESHelper(event)

    try:
        esh.validate_event()
    except AssertionError as err:
        return esh.return_code(message=err, code=500)


    fetch = esh.event()['body']['elements'].get('fetch')

    if not fetch or not fetch.get('conditions'):
        return esh.return_code(message='Wrong or Missing Information', code=404)

    if 'user' == fetch.get('conditions').get('type'):
        location = fetch.get('location')
        user_id = fetch['conditions']['user']
        location = fetch['conditions'].get('location')

        user_hit = None
        try:
            user_hit = esh.es().get(index='users', id=user_id, doc_type='_doc')
        except Exception as err:
            print('Could not fetch user')

        if not user_hit or not user_hit['found']:
            return esh.return_code(message='User does not exist. Please contact support', code=400)

        # eshtracking = UtilitiesTracking(esh=esh, user_id=user_id, user=user_hit['_source'], location=location, active=True)
            
        # try:
        #     eshtracking.update_tracking()
        # except AssertionError as err:
        #     print(err)

        return esh.return_code_with(code=200, message='OK!', elements={'user': user_hit['_source']})
   
    elif 'post' == fetch.get('conditions').get('type'):

        location = fetch['conditions'].get('location')

        if not location and (cfilter == 'new' or cfilter == 'hot' or cfilter == 'top'):
            return esh.return_code(message='Location must be enabled', code=400)

        if cfilter == 'new':

        elif cfilter == 'hot':
        
        elif cfilter == 'top':
      
        elif cfilter == 'explore':
            
        elif cfilter == 'personal':
            if csubfilter == 'reply':

            elif csubfilter == 'post':

            elif csubfilter == 'vote':

        elif cfilter == 'keyword':
        elif cfilter == 'location':



        return esh.return_code_with(code=200, message='OK!', elements={'post': user_hit['_source']})
        