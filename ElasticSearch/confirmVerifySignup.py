# -*- coding: utf-8 -*-


from eshelper import ESHelper, UtilitiesTracking, TemplateUser


def lambda_handler(event, context):

    esh = ESHelper(event)

    event['response']['autoConfirmUser'] = True

    # Set the phone number as verified if it is in the request
    if 'phone_number' in event['request']['userAttributes']:
        event['response']['autoVerifyPhone'] = True

    user_id = event['request']['userAttributes']['phone_number']
    country_code = event['request']['userAttributes']['custom:countryCode']
    password = event['request']['userAttributes']['custom:password']

    eshuser = TemplateUser(esh=esh, user_id=user_id, password=password, country_code=country_code)

    user_item = eshuser.blank_user()

    try:
        esh.es().create(index= 'users', id=user_id, doc_type='_doc', body=user_item)
    except Exception as err:
        print(err)
        raise ValueError('We\'re a little busy at the moment. Please try again in a little.')

    eshtracking = UtilitiesTracking(esh=esh, user_id=user_id, user=user_item, active=True, registration=True)
        
    try:
        eshtracking.update_tracking()
    except AssertionError as err:
        print(err)
    
    # Return to Amazon Cognito
    return event