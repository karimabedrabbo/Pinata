# -*- coding: utf-8 -*-

from eshelper import ESHelper, UtilitiesTracking

def lambda_handler(event, context):

	esh = ESHelper(event)

	try:
		esh.validate_event()
	except AssertionError as err:
		return esh.return_code(message=err, code=500)


	fetch = esh.event()['body']['elements'].get('fetch')

	if not fetch or not fetch.get('conditions') or not 'user' == fetch.get('conditions').get('type') or not fetch.get('conditions').get('user'):
		return esh.return_code(message='Wrong or Missing Information', code=404)

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

	eshtracking = UtilitiesTracking(esh=esh, user_id=user_id, user=user_hit['_source'], location=location, active=True)
		
	try:
		eshtracking.update_tracking()
	except AssertionError as err:
		print(err)

	return esh.return_code_with(code=200, message='OK!', elements={'user': user_hit['_source']})
	