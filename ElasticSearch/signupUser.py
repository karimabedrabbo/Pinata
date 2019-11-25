# -*- coding: utf-8 -*-

from eshelper import ESHelper, UtilitiesTracking, TemplateUser


def lambda_handler(event, context):

	esh = ESHelper(event)

	try:
		esh.validate_event()
	except AssertionError as err:
		return esh.return_code(message=err, code=500)

	signup = esh.event()['body']['elements'].get('signup')

	if not signup or not signup.get('userId') or not signup.get('countryCode') or not signup.get('password'):
		return esh.return_code(message='Missing Information', code=404)

	location = signup.get('location')
	user_id = signup['userId']
	country_code = signup['countryCode']
	password = signup['password']

	if esh.es().exists(index='users', doc_type='_doc', id=user_id):
		return esh.return_code(message='User Already Exists', code=404)

	eshuser = TemplateUser(esh=esh, user_id=user_id, password=password, country_code=country_code)

	user_item = eshuser.blank_user()

	try:
		esh.es().create(index= 'users', id=user_id, doc_type='_doc', body=user_item)
	except Exception as err:
		print(err)
		return esh.return_code(message='Could not put user', code=400)

	eshtracking = UtilitiesTracking(esh=esh, user_id=user_id, user=user_item, location=location, active=True, registration=True)
		
	try:
		eshtracking.update_tracking()
	except AssertionError as err:
		print(err)
	
	return esh.return_code(message='OK!', code=200)