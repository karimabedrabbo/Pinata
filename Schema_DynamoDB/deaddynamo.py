#SIGNUP FILE

#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#DEAD DYNAMO
#################################################################################################
#IMPORTS

# from boto3.dynamodb.types import TypeDeserializer, TypeSerializer
# from decimal import Decimal
# from boto3.dynamodb.conditions import Key, Attr
# from botocore.errorfactory import ClientError


#################################################################################################
#RELEVANT STATEMENTS

# db = session.resource('dynamodb')
# user_table = db.Table('User')
# rint = randint(0, 29)

#################################################################################################
#USER CHECKING

# user_response = user_table.get_item(
# 		Key={
# 			'uId': signup.get('user_id')
# 		}
# 	)

# 	if user_response.get('Item'):
# 		return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'User Already Exists'})}


#################################################################################################
#USER ADDING

# user_item = {
# 			'uId': signup.get('user_id'),
# 			'cTp': created_at_milliseconds,
# 			'mTp': created_at_milliseconds,
# 			'ctryCd': signup.get('country_code'),
# 			'usrnm': new_user_handle,
# 			'krmaSr': 0,
# 			'pswd': signup.get('password'),
# 			'acsCd': 'usr',
# 			'rnd': rint,
# 			'uvrf': 1
# 		}

# try:
# 	user_table.put_item(
# 		Item=user_item)
# except ClientError as err:
# 	print("Could not put user item")
# 	print(err)
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Service too busy. Please try again.'})}

#################################################################################################

#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

#################################################################################################
#TABLE INFO
#################################################################################################
#User Table:

#Key: uId
#c: bnEnd (ban end timestamp)
#d: ctryCd (country code)
#e: usrnm (username)
#f: phnNb (phone number)
#g: krmaSr (karma score)
#h: cTp (created timestamp)
#h: mTp (modified timestamp)
#i: pswd (password)
#j: pmL16Id (permanent location level 16)
#k: pmLng (permanent longitude)
#l: pmLat (permanent latitude)
#j: pmCtyNm (permanent city name)
#j: pmCmpNm (permanent campus name)
#j: pmCmpAl (permanent campus alias)
#n: bio (biography)
#o: bDy (birthday)
#o: nme (name)
#o: pfBkKy (profile picture bucket key)
#p: acsCd (options are USR, MOD, ADM with random integer 00-30)
#q: uOpts (user options: map)
#r: uRpl (user replies list)
#g: uPst (user posts list)
#a: uRpst (user reposts list)
#f: uStk (user stickers list)
#g: lFlwCt (locations following count)
#g: uFrdCt (users friend count)
#h: uFrd (users friends list)
#h: uFrdRqSt (users friend request sent)
#h: uFrdRqRc (users friend request recieved)
#g: lFlw (locations following list)
#f: uBlk (user blocking list)
#g: uTph (user trophies list)
#g: rnd (0-30)
#g: vrf (0 or 1)

#-----------------------------------------------------------------------------------------------
#GSI1: 
#description: retrieve user by phone number
#key: phnNb
#-----------------------------------------------------------------------------------------------
#GSI2: 
#description: retrieve user by user name
#key: rnd
#range: usrnm

#################################################################################################

#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$


#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#DEAD DYNAMO

#################################################################################################
#IMPORTS

# from boto3.dynamodb.types import TypeDeserializer, TypeSerializer
# import decimal
# from random import randint
# from boto3.dynamodb.conditions import Key, Attr
# from botocore.errorfactory import ClientError
# import hmac
# import hashlib
# import base64
# import uuid


#################################################################################################
#HELPER FUNCTIONS

# class DecimalEncoder(json.JSONEncoder):
# 	def default(self, o):
# 		if isinstance(o, decimal.Decimal):
# 			if o % 1 > 0:
# 				return float(o)
# 			else:
# 				return int(o)
# 		return super(DecimalEncoder, self).default(o)

# def user_dne_table(key, key_is_phone=True):
# 	session = boto3.Session()
# 	db = session.resource('dynamodb')
# 	user_table = db.Table('User')
# 	cog_client = boto3.client('cognito-idp')
# 	rint = randint(0, 29)
# 	new_user = {}
# 	if not key_is_phone:
# 		# new_user['user_id'] = key
# 		#Key is user id so search cognito for sub
# 		cog_response = cog_client.list_users(
# 			UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
# 			AttributesToGet=[
# 				'sub',
# 			],
# 			Limit=1,
# 			Filter='sub = ' + '\"' + key + '\"' 
# 		)
# 	else:
# 		#Key is phone so retrieve user from cognito
# 		new_user['phone_number'] = key
# 		cog_response = cog_client.admin_get_user(
# 			UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
# 			Username=key
# 		)
# 		# to_retrieve_id = cog_client.list_users(
# 		#     UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
# 		#     AttributesToGet=[
# 		#         'phone_number', 'sub'
# 		#     ],
# 		#     Limit=1,
# 		#     Filter='phone_number = ' + '\"' + key + '\"' 
# 		# )
# 		# for attr in to_retrieve_id.get('Users')[0].get('Attributes'):
# 		# 	if attr['Name'] == 'sub':
# 		# 		new_user['user_id'] = attr['Value']
# 	if cog_response:
# 		for attr in cog_response.get('UserAttributes'):
# 			if attr['Name'] == 'custom:ctryCd':
# 				new_user['country_code'] = attr['Value']
# 			elif attr['Name'] == 'custom:usrnm':
# 				new_user['user_handle'] = attr['Value']
# 			elif attr['Name'] == 'custom:uvrf':
# 				new_user['verification'] = attr['Value']
# 			elif attr['Name'] == 'custom:acsCd':
# 				new_user['access_code'] = attr['Value']
# 			elif attr['Name'] == 'custom:phone_number':
# 				new_user['phone_number'] = attr['Value']
# 	else:
# 		return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not retrieve user from Cognito when user not in table on retrieval call.'})}


# 	if new_user.get('country_code') and new_user.get('user_handle') and new_user.get('verification') and new_user.get('access_code') and new_user.get('phone_number'):
# 		new_user['password'] = str(uuid.uuid4())

# 		# secret_to_encode = new_user['phone_number'] + os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID']
# 		# encoded_secret = hmac.new(bytes(os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID'], 'utf8'), msg=secret_to_encode.encode('utf-8'),
# 		# 			   digestmod=hashlib.sha256).digest()
# 		# encoded_secret_base64 = base64.b64encode(encoded_secret).decode()

# 		_response = cog_client.admin_delete_user(
# 				UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
# 				Username=new_user.get('phone_number')
# 			)
# 		new_user_in_cognito = cog_client.sign_up(
# 			ClientId=os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID_WEB'],
# 			# SecretHash=encoded_secret_base64,
# 			Username=new_user['phone_number'],
# 			Password=new_user['password'],
# 			UserAttributes=[
# 				{
# 					'Name': 'custom:ctryCd',
# 					'Value': new_user['country_code']
# 				},
# 				{
# 					'Name': 'custom:uvrf',
# 					'Value': new_user['verification']
# 				},
# 				{
# 					'Name': 'custom:usrnm',
# 					'Value': new_user['user_handle']
# 				},
# 				{
# 					'Name': 'custom:acsCd',
# 					'Value': new_user['access_code']
# 				},
# 			],
# 		)
		
# 		new_user['user_id'] = new_user_in_cognito.get('UserSub')
# 		if not new_user.get('user_id'):
# 			print('Could not create user after deleting user')
# 			return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not find sub attr after deleting and reinstantiating user'})} 
# 		try:
# 			new_user_put = {
# 			'uId': new_user['user_id'],
# 			'cTp': int(time.time()),
# 			'ctryCd': new_user['country_code'],
# 			'usrnm': new_user['user_handle'],
# 			'phnNb': new_user['phone_number'],
# 			'krmaSr': 0,
# 			'pswd': new_user['password'],
# 			'acsCd': new_user['access_code'],
# 			'rnd': rint,
# 			'uvrf': rint
# 			}
# 			user_table.put_item(
# 				Item=new_user_put)
# 			return new_user_put
# 		except ClientError as err:
# 			print("Could not put user item")
# 			print(err.response)
# 			return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not put user on retrieval call'})}
# 	else:
# 		return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'User retrieved from Cognito did not have sufficient attributes when user not in table on retrieval call.'})}

# if phone_number and not user_id:
# 			try:
# 				responsePhoneQuery = user_table.query(
# 				IndexName='phnNb-index',
# 				KeyConditionExpression=Key('phnNb').eq(phone_number)
# 				)
# 				itemsPhoneQuery = responsePhoneQuery.get('Items')
# 			except ClientError as err:
# 				return {'statusCode': 500, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 500, 'message': 'Internal Error. Could not query by phone number.'})}
# 			if len(itemsPhoneQuery) > 1:
# 				return {'statusCode': 409, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 409, 'message': 'Conflicting phone numbers'})}
# 			elif not itemsPhoneQuery:
# 				potential_put = user_dne_table(phone_number, key_is_phone=True)
# 				if potential_put.get('statusCode'):
# 					return potential_put
# 				else:
# 					itemsPhoneQuery.append(potential_put)

# 			new_user = itemsPhoneQuery[-1]
# 			new_user = json.loads(json.dumps(new_user, cls=DecimalEncoder))
# 			return {'statusCode': 200, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 200, 'message': 'OK!', 'user': new_user})}

# if user_id:
# 			try:
# 				new_user = user_table.get_item(
# 					Key={
# 						'uId': user_id
# 					}
# 					)
# 			except ClientError as err:
# 				return {'statusCode': 500, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 500, 'message': 'Internal Error. Could not query by phone number.'})}
# 			else:
# 				if not new_user:
# 					potential_put = user_dne_table(user_id, key_is_phone=False)
# 					if potential_put.get('statusCode'):
# 						return potential_put
# 					else:
# 						new_user = potential_put
# 				return {'statusCode': 200, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 200, 'message': 'OK!', 'user': new_user})}

#################################################################################################
#RETRIEVING USER

# try:
# 	user_response = user_table.get_item(
# 				Key={
# 					'uId': user_id
# 				}
# 				)
# except ClientError as err:
# 	print(err)
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Service is too busy. Please try again.'})}

# if not user_response.get('Item'):
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 422, 'message': 'User does not exist. Please contact support.'})}

# retrieved_user_final = json.loads(json.dumps(user_response.get('Item'), cls=DecimalEncoder))
# return {'statusCode': 200, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 200, 'message': 'OK!', 'user': retrieved_user_final})}

#################################################################################################

#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

#RETRIEVE USER FILE


#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#DEAD DYNAMO

#################################################################################################
#IMPORTS

# from boto3.dynamodb.types import TypeDeserializer, TypeSerializer
# import decimal
# from random import randint
# from boto3.dynamodb.conditions import Key, Attr
# from botocore.errorfactory import ClientError
# import hmac
# import hashlib
# import base64
# import uuid


#################################################################################################
#HELPER FUNCTIONS

# class DecimalEncoder(json.JSONEncoder):
# 	def default(self, o):
# 		if isinstance(o, decimal.Decimal):
# 			if o % 1 > 0:
# 				return float(o)
# 			else:
# 				return int(o)
# 		return super(DecimalEncoder, self).default(o)

# def user_dne_table(key, key_is_phone=True):
# 	session = boto3.Session()
# 	db = session.resource('dynamodb')
# 	user_table = db.Table('User')
# 	cog_client = boto3.client('cognito-idp')
# 	rint = randint(0, 29)
# 	new_user = {}
# 	if not key_is_phone:
# 		# new_user['user_id'] = key
# 		#Key is user id so search cognito for sub
# 		cog_response = cog_client.list_users(
# 			UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
# 			AttributesToGet=[
# 				'sub',
# 			],
# 			Limit=1,
# 			Filter='sub = ' + '\"' + key + '\"' 
# 		)
# 	else:
# 		#Key is phone so retrieve user from cognito
# 		new_user['phone_number'] = key
# 		cog_response = cog_client.admin_get_user(
# 			UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
# 			Username=key
# 		)
# 		# to_retrieve_id = cog_client.list_users(
# 		#     UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
# 		#     AttributesToGet=[
# 		#         'phone_number', 'sub'
# 		#     ],
# 		#     Limit=1,
# 		#     Filter='phone_number = ' + '\"' + key + '\"' 
# 		# )
# 		# for attr in to_retrieve_id.get('Users')[0].get('Attributes'):
# 		# 	if attr['Name'] == 'sub':
# 		# 		new_user['user_id'] = attr['Value']
# 	if cog_response:
# 		for attr in cog_response.get('UserAttributes'):
# 			if attr['Name'] == 'custom:ctryCd':
# 				new_user['country_code'] = attr['Value']
# 			elif attr['Name'] == 'custom:usrnm':
# 				new_user['user_handle'] = attr['Value']
# 			elif attr['Name'] == 'custom:uvrf':
# 				new_user['verification'] = attr['Value']
# 			elif attr['Name'] == 'custom:acsCd':
# 				new_user['access_code'] = attr['Value']
# 			elif attr['Name'] == 'custom:phone_number':
# 				new_user['phone_number'] = attr['Value']
# 	else:
# 		return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not retrieve user from Cognito when user not in table on retrieval call.'})}


# 	if new_user.get('country_code') and new_user.get('user_handle') and new_user.get('verification') and new_user.get('access_code') and new_user.get('phone_number'):
# 		new_user['password'] = str(uuid.uuid4())

# 		# secret_to_encode = new_user['phone_number'] + os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID']
# 		# encoded_secret = hmac.new(bytes(os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID'], 'utf8'), msg=secret_to_encode.encode('utf-8'),
# 		# 			   digestmod=hashlib.sha256).digest()
# 		# encoded_secret_base64 = base64.b64encode(encoded_secret).decode()

# 		_response = cog_client.admin_delete_user(
# 				UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
# 				Username=new_user.get('phone_number')
# 			)
# 		new_user_in_cognito = cog_client.sign_up(
# 			ClientId=os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID_WEB'],
# 			# SecretHash=encoded_secret_base64,
# 			Username=new_user['phone_number'],
# 			Password=new_user['password'],
# 			UserAttributes=[
# 				{
# 					'Name': 'custom:ctryCd',
# 					'Value': new_user['country_code']
# 				},
# 				{
# 					'Name': 'custom:uvrf',
# 					'Value': new_user['verification']
# 				},
# 				{
# 					'Name': 'custom:usrnm',
# 					'Value': new_user['user_handle']
# 				},
# 				{
# 					'Name': 'custom:acsCd',
# 					'Value': new_user['access_code']
# 				},
# 			],
# 		)
		
# 		new_user['user_id'] = new_user_in_cognito.get('UserSub')
# 		if not new_user.get('user_id'):
# 			print('Could not create user after deleting user')
# 			return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not find sub attr after deleting and reinstantiating user'})} 
# 		try:
# 			new_user_put = {
# 			'uId': new_user['user_id'],
# 			'cTp': int(time.time()),
# 			'ctryCd': new_user['country_code'],
# 			'usrnm': new_user['user_handle'],
# 			'phnNb': new_user['phone_number'],
# 			'krmaSr': 0,
# 			'pswd': new_user['password'],
# 			'acsCd': new_user['access_code'],
# 			'rnd': rint,
# 			'uvrf': rint
# 			}
# 			user_table.put_item(
# 				Item=new_user_put)
# 			return new_user_put
# 		except ClientError as err:
# 			print("Could not put user item")
# 			print(err.response)
# 			return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not put user on retrieval call'})}
# 	else:
# 		return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'User retrieved from Cognito did not have sufficient attributes when user not in table on retrieval call.'})}

# if phone_number and not user_id:
# 			try:
# 				responsePhoneQuery = user_table.query(
# 				IndexName='phnNb-index',
# 				KeyConditionExpression=Key('phnNb').eq(phone_number)
# 				)
# 				itemsPhoneQuery = responsePhoneQuery.get('Items')
# 			except ClientError as err:
# 				return {'statusCode': 500, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 500, 'message': 'Internal Error. Could not query by phone number.'})}
# 			if len(itemsPhoneQuery) > 1:
# 				return {'statusCode': 409, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 409, 'message': 'Conflicting phone numbers'})}
# 			elif not itemsPhoneQuery:
# 				potential_put = user_dne_table(phone_number, key_is_phone=True)
# 				if potential_put.get('statusCode'):
# 					return potential_put
# 				else:
# 					itemsPhoneQuery.append(potential_put)

# 			new_user = itemsPhoneQuery[-1]
# 			new_user = json.loads(json.dumps(new_user, cls=DecimalEncoder))
# 			return {'statusCode': 200, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 200, 'message': 'OK!', 'user': new_user})}

# if user_id:
# 			try:
# 				new_user = user_table.get_item(
# 					Key={
# 						'uId': user_id
# 					}
# 					)
# 			except ClientError as err:
# 				return {'statusCode': 500, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 500, 'message': 'Internal Error. Could not query by phone number.'})}
# 			else:
# 				if not new_user:
# 					potential_put = user_dne_table(user_id, key_is_phone=False)
# 					if potential_put.get('statusCode'):
# 						return potential_put
# 					else:
# 						new_user = potential_put
# 				return {'statusCode': 200, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 200, 'message': 'OK!', 'user': new_user})}

#################################################################################################
#RETRIEVING USER

# try:
# 	user_response = user_table.get_item(
# 				Key={
# 					'uId': user_id
# 				}
# 				)
# except ClientError as err:
# 	print(err)
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Service is too busy. Please try again.'})}

# if not user_response.get('Item'):
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 422, 'message': 'User does not exist. Please contact support.'})}

# retrieved_user_final = json.loads(json.dumps(user_response.get('Item'), cls=DecimalEncoder))
# return {'statusCode': 200, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 200, 'message': 'OK!', 'user': retrieved_user_final})}

#################################################################################################

#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

#################################################################################################
#TABLE INFO
#################################################################################################

#User Table:
#Key: uId

#c: bnEnd (ban end timestamp)
#d: ctryCd (country code)
#e: usrnm (username)
#f: phnNb (phone number)
#g: krmaSr (karma score)
#h: cTp (created timestamp)
#h: mTp (modified timestamp)
#i: pswd (password)
#j: pmL16Id (permanent location level 16)
#k: pmLng (permanent longitude)
#l: pmLat (permanent latitude)
#j: pmCtyNm (permanent city name)
#j: pmCmpNm (permanent campus name)
#j: pmCmpAl (permanent campus alias)
#n: bio (biography)
#o: bDy (birthday)
#o: nme (name)
#o: pfBkKy (profile picture bucket key)
#p: acsCdRInt (options are USR, MOD, ADM with random integer 00-30)
#q: uOpts (user options: map)
#r: uRpl (user replies list)
#g: uPst (user posts list)
#a: uRpst (user reposts list)
#f: uStk (user stickers list)
#g: lFlwCt (locations following count)
#g: uFrdCt (users friend count)
#h: uFrd (users friends list)
#h: uFrdRqSt (users friend request sent)
#h: uFrdRqRc (users friend request recieved)
#g: lFlw (locations following list)
#f: uBlk (user blocking list)
#g: uTph (user trophies list)
#g: rnd (0-30)
#g: vrf (0 or 1)

#-----------------------------------------------------------------------------------------------
#GSI1: 
#description: retrieve user by phone number
#key: phnNb
#-----------------------------------------------------------------------------------------------
#GSI2: 
#description: retrieve user by user name
#key: rnd
#range: usrnm

#################################################################################################

#SIGNIN USER FILE
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#DEAD DYNAMO

#################################################################################################
#IMPORTS

# from boto3.dynamodb.types import TypeDeserializer, TypeSerializer
# from decimal import Decimal
# import hmac
# import hashlib
# import base64
# from boto3.dynamodb.conditions import Key, Attr
# from botocore.errorfactory import ClientError
# from random import randint

#################################################################################################
#RELEVANT STATEMENTS

# db = session.resource('dynamodb')
# user_table = db.Table('User')

#################################################################################################
#HELPER FUNCTIONS

# def user_dne_table(key, key_is_phone=True):
#     session = boto3.Session()
#     db = session.resource('dynamodb')
#     user_table = db.Table('User')
#     cog_client = boto3.client('cognito-idp')
#     rint = randint(0, 29)
#     new_user = {}
#     if not key_is_phone:
#         # new_user['user_id'] = key
#         #Key is user id so search cognito for sub
#         cog_response = cog_client.list_users(
#             UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
#             AttributesToGet=[
#                 'sub',
#             ],
#             Limit=1,
#             Filter='sub = ' + '\"' + key + '\"' 
#         )
#     else:
#         #Key is phone so retrieve user from cognito
#         new_user['phone_number'] = key
#         cog_response = cog_client.admin_get_user(
#             UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
#             Username=key
#         )
#         # to_retrieve_id = cog_client.list_users(
#         #     UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
#         #     AttributesToGet=[
#         #         'phone_number', 'sub'
#         #     ],
#         #     Limit=1,
#         #     Filter='phone_number = ' + '\"' + key + '\"' 
#         # )
#         # for attr in to_retrieve_id.get('Users')[0].get('Attributes'):
#         #   if attr['Name'] == 'sub':
#         #       new_user['user_id'] = attr['Value']
#     if cog_response:
#         for attr in cog_response.get('UserAttributes'):
#             if attr['Name'] == 'custom:ctryCd':
#                 new_user['country_code'] = attr['Value']
#             elif attr['Name'] == 'custom:usrnm':
#                 new_user['user_handle'] = attr['Value']
#             elif attr['Name'] == 'custom:uvrf':
#                 new_user['verification'] = attr['Value']
#             elif attr['Name'] == 'custom:acsCd':
#                 new_user['access_code'] = attr['Value']
#             elif attr['Name'] == 'custom:phone_number':
#                 new_user['phone_number'] = attr['Value']
#     else:
#         return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not retrieve user from Cognito when user not in table on retrieval call.'})}


#     if new_user.get('country_code') and new_user.get('user_handle') and new_user.get('verification') and new_user.get('access_code') and new_user.get('phone_number'):
#         new_user['password'] = str(uuid.uuid4())

#         # secret_to_encode = new_user['phone_number'] + os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID']
#         # encoded_secret = hmac.new(bytes(os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID'], 'utf8'), msg=secret_to_encode.encode('utf-8'),
#         #              digestmod=hashlib.sha256).digest()
#         # encoded_secret_base64 = base64.b64encode(encoded_secret).decode()

#         _response = cog_client.admin_delete_user(
#                 UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
#                 Username=new_user.get('phone_number')
#             )
#         new_user_in_cognito = cog_client.sign_up(
#             ClientId=os.environ['MOBILE_HUB_COGNITO_IDENTITY_USER_POOL_APP_CLIENT_ID_WEB'],
#             # SecretHash=encoded_secret_base64,
#             Username=new_user['phone_number'],
#             Password=new_user['password'],
#             UserAttributes=[
#                 {
#                     'Name': 'custom:ctryCd',
#                     'Value': new_user['country_code']
#                 },
#                 {
#                     'Name': 'custom:uvrf',
#                     'Value': new_user['verification']
#                 },
#                 {
#                     'Name': 'custom:usrnm',
#                     'Value': new_user['user_handle']
#                 },
#                 {
#                     'Name': 'custom:acsCd',
#                     'Value': new_user['access_code']
#                 },
#             ],
#         )
		
#         new_user['user_id'] = new_user_in_cognito.get('UserSub')
#         if not new_user.get('user_id'):
#             print('Could not create user after deleting user')
#             return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not find sub attr after deleting and reinstantiating user'})} 
#         try:
#             new_user_put = {
#             'uId': new_user['user_id'],
#             'cTp': int(time.time()),
#             'ctryCd': new_user['country_code'],
#             'usrnm': new_user['user_handle'],
#             'phnNb': new_user['phone_number'],
#             'krmaSr': 0,
#             'pswd': new_user['password'],
#             'acsCd': new_user['access_code'],
#             'rnd': rint,
#             'uvrf': rint
#             }
#             user_table.put_item(
#                 Item=new_user_put)
#             return new_user_put
#         except ClientError as err:
#             print("Could not put user item")
#             print(err.response)
#             return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Could not put user on retrieval call'})}
#     else:
#         return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'User retrieved from Cognito did not have sufficient attributes when user not in table on retrieval call.'})}



		# if user_id and not phone_number:
		#     to_retrieve_id = cog_client.list_users(
		#     UserPoolId=os.environ['MOBILE_HUB_COGNITO_USER_POOL_ID'],
		#     AttributesToGet=[
		#         'phone_number', 'sub'
		#     ],
		#     Limit=1,
		#     Filter='sub = ' + '\"' + user_id + '\"' 
		#     )
		#     for attr in to_retrieve_id.get('Users')[0].get('Attributes'):
		#       if attr['Name'] == 'sub':
		#           phone_number = attr['Value']

# if phone_number and not user_id:
#             responsePhoneQuery = user_table.query(
#             IndexName='phnNb-index',
#             KeyConditionExpression=Key('phnNb').eq(phone_number)
#             )
#             itemsPhoneQuery = responsePhoneQuery.get('Items')
#             if not itemsPhoneQuery:
#                 potential_put = user_dne_table(phone_number, key_is_phone=True)
#                 if potential_put.get('statusCode'):
#                     return potential_put
#                 else:
#                     itemsPhoneQuery.append(potential_put)
#             elif len(itemsPhoneQuery) > 1:
#                 return {'statusCode': 409, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 409, 'message': 'Conflicting phone numbers'})}
#             user_id = itemsPhoneQuery[-1]['uId']

#################################################################################################
#CHECK USER EXISTS

# if not user_id:
# 	return {'statusCode': 422, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 422, 'message': 'Missing Information.'})}

# user_response = user_table.get_item(
# 	Key={
# 		'uId': user_id
# 	}
# )

# if not user_response.get('Item'):
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'User does not exist. Please contact the support team.'})}

#################################################################################################
#UPDATE VERIFIED ATTRIBUTE (2 VERSIONS)

# try: 
# 	user_table.update_item(
# 		Key={
# 				'uId': user_id,
# 			},
# 		UpdateExpression='SET uvrf = :u',
# 		ExpressionAttributeValues={
# 					':u': 0
# 			}
# 		)
# except ClientError as err:
# 	print('Unverified Attribute does not exist.')



# try: 
# 	user_table.update_item(
# 		Key={
# 				'uId': user_id,
# 			},
# 		UpdateExpression='REMOVE uvrf',
# 		ConditionExpression='attribute_exists(uvrf)',
# 		)
# except ClientError as err:
# 	print('Nothing to update. Perfectly fine.')


#################################################################################################

#User Table:
#Key: uId

#c: bnEnd (ban end timestamp)
#d: ctryCd (country code)
#e: usrnm (username)
#f: phnNb (phone number)
#g: krmaSr (karma score)
#h: cTp (created timestamp)
#h: mTp (modified timestamp)
#i: pswd (password)
#j: pmL16Id (permanent location level 16)
#k: pmLng (permanent longitude)
#l: pmLat (permanent latitude)
#j: pmCtyNm (permanent city name)
#j: pmCmpNm (permanent campus name)
#j: pmCmpAl (permanent campus alias)
#n: bio (biography)
#o: bDy (birthday)
#o: nme (name)
#o: pfBkKy (profile picture bucket key)
#p: acsCdRInt (options are USR, MOD, ADM with random integer 00-30)
#q: uOpts (user options: map)
#r: uRpl (user replies list)
#g: uPst (user posts list)
#a: uRpst (user reposts list)
#f: uStk (user stickers list)
#g: lFlwCt (locations following count)
#g: uFrdCt (users friend count)
#h: uFrd (users friends list)
#h: uFrdRqSt (users friend request sent)
#h: uFrdRqRc (users friend request recieved)
#g: lFlw (locations following list)
#f: uBlk (user blocking list)
#g: uTph (user trophies list)
#g: rnd (0-30)
#g: vrf (0 or 1)

#-----------------------------------------------------------------------------------------------
#GSI1: 
#description: retrieve user by phone number
#key: phnNb
#-----------------------------------------------------------------------------------------------
#GSI2: 
#description: retrieve user by user name
#key: rnd
#range: usrnm

#################################################################################################




#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#DEAD DYNAMO

#################################################################################################
#IMPORTS
# from random import randint
# from boto3.dynamodb.types import TypeDeserializer, TypeSerializer
# import hashlib
# from collections import Sequence, Set
# from decimal import Decimal
# import backoff
# from boto3.dynamodb.conditions import Key, Attr
# from botocore.errorfactory import ClientError

#################################################################################################
#HELPER FUNCTIONS

# class DecimalEncoder(json.JSONEncoder):
# 	def default(self, o):
# 		if isinstance(o, Decimal):
# 			if o % 1 > 0:
# 				return float(o)
# 			else:
# 				return int(o)
# 		return super(DecimalEncoder, self).default(o)

# def _dynamo_sanitize(data):
# 	""" Sanitizes an object so it can be updated to dynamodb (recursive) """
# 	if not data and isinstance(data, (str, Set)):
# 		new_data = None  # empty strings/sets are forbidden by dynamodb
# 	elif isinstance(data, (str, bool)):
# 		new_data = data  # important to handle these one before sequence and int!
# 	elif isinstance(data, dict):
# 		new_data = {key: _dynamo_sanitize(data[key]) for key in data}
# 	elif isinstance(data, Sequence):
# 		new_data = [_dynamo_sanitize(item) for item in data]
# 	elif isinstance(data, Set):
# 		new_data = {_dynamo_sanitize(item) for item in data}
# 	elif isinstance(data, (float, complex)):
# 		new_data = Decimal(str(data))
# 	else:
# 		new_data = data
# 	return new_data


# def location_hashed(s):
# 	onebyte = (hashlib.md5(s.encode('utf8')).digest()[0])
# 	return onebyte % 30

# @backoff.on_predicate(backoff.full_jitter, lambda response: response.get('UnprocessedItems') != None, max_time=30, value=30)
# def batch_input(table, batch_list):
# 	try:
# 		#incomplete
# 		response = table.batch_write_item(list(map(lambda x: _dynamo_sanitize(x), batch_list)))
# 		return response
# 	except ClientError as err:
# 		print(err)
# 		return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Event Local table too busy'})}


#################################################################################################
#CELL DENSITY

# 	try:
# 		location_hour_table.update_item(
# 			Key={
# 				'lId': level30.parent(i).to_token(),
# 				'hrTp': created_at_hour.strftime('%Y-%m-%d-%H')
# 			},
# 			UpdateExpression='ADD dsy :v SET eTp = :e',
# 			ExpressionAttributeValues={
# 				':v': 1,
# 				':e': expiring_at
# 			},
# 		)
# 	except ClientError as err:
# 		print(err)

# 	try:
# 		location_day_table.update_item(
# 			Key={
# 				'lId': level30.parent(i).to_token(),
# 				'dyTp': created_at_day.strftime('%Y-%m-%d-%H')
# 			},
# 			UpdateExpression='ADD dsy :v SET eTp = :e',
# 			ExpressionAttributeValues={
# 				':v': 1,
# 				':e': expiring_at
# 			},
# 		)
# 	except ClientError as err:
# 		print(err)

# 	try:
# 		location_week_table.update_item(
# 			Key={
# 				'lId': level30.parent(i).to_token(),
# 				'wkTp': created_at_week.strftime('%Y-%m-%d-%H')
# 			},
# 			UpdateExpression='ADD dsy :v SET eTp = :e',
# 			ExpressionAttributeValues={
# 				':v': 1,
# 				':e': expiring_at
# 			},
# 		)
# 	except ClientError as err:
# 		print(err)

# level10_neighbors = list(map(lambda x: x.to_token(), list(level30.parent(10).get_all_neighbors(10)) + [level30.parent(10)]))

# for token in level10_neighbors:
# 	try:
# 		user_location_table_item_put = {
# 			'lId': token,
# 			'uId': author,
# 		}
# 		user_location_table.put_item(Item=_dynamo_sanitize(user_location_table_item_put))
# 	except ClientError as err:
# 		print(err)


#################################################################################################
#EVENT LOCAL

# for i in range(10,17):
# 	event_local_item_put = {
# 		'eCdLId': event_code + '_' + level30.parent(i).to_token(),
# 		'cTp': created_at_milliseconds,
# 		'eTp': expiring_at,
# 		'pId': pId,
# 		'uId': author,
# 		'vlSr': initial_rank,
# 	}

# 	try:
# 		event_local_table.put_item(Item=_dynamo_sanitize(event_local_item_put))
# 	except ClientError as err:
# 		print(err)

#################################################################################################
#CHECK POPULATION FOR INITIAL RANKING SCORE

# try:
# 	location_week_table_response = location_week_table.get_item(
# 		Key={
# 			'lId': level30.parent(10).to_token(),
# 			'wkTp': created_at_week.strftime('%Y-%m-%d-%H')
# 		}
# 	)
# except ClientError as err:
# 	print(err)
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Location table too busy to query'})}

# if not location_week_table_response.get('Item'):
# 	location_week_table_response['Item'] = {'dsy': 0}


#initial_rank = value_ranking_population(location_week_table_response.get('Item').get('dsy',0), 0, 0, created=created_at_milliseconds // (1000 * 1000), now=created_at_milliseconds // (1000 * 1000))


#################################################################################################
#POSTING

# lhsh = location_hashed(level30.parent(10).to_token())

# post_item_put = {'pId': pId,
# 				'cTp': created_at_milliseconds,
# 				'eCd': event_code + '_' + str(lhsh),
# 				'vlSr': initial_rank,
# 				'pkVlSr': initial_rank,
# 				'acsCd': 'rvw',
# 				'rplSr': 0,
# 				'shrSr': 0,
# 				'vtSr': 0,
# 				'dwnVt': 0,
# 				'upVt': 0,
# 				'rpySr': 0,
# 				'uId': author,
# 				'atchm': attachments,
# 				'mta': meta,
# 				'chld': children
# 			}

# post_table.put_item(
# 	Item=_dynamo_sanitize(post_item_put)
# )

#################################################################################################
#ACTIVITY ALL TABLE

# try:
# 	activity_all_item_put = {
# 		'eCd': event_code + '_' + str(lhsh),
# 		'cTp': created_at_milliseconds,
# 		'eTp': expiring_at,
# 		'pId': pId,
# 		'uId': author,
# 		'vlSr': initial_rank,
		
# 	}
# 	activity_all_table.put_item(Item=_dynamo_sanitize(activity_all_item_put))
# except ClientError as err:
# 	print(err)
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Activity all table too busy'})}

#################################################################################################
#APPEND USER POST TO USER TABLE


# try:
# 	user_table.update_item(
# 	Key={
# 		'uId': author
# 	},
# 	UpdateExpression='SET uPst = list_append(:pId, uPst)',
# 	ExpressionAttributeValues={
# 		':pId': [pId]
# 	},
# 	ConditionExpression="attribute_exists(uPst)",
# )
# except ClientError as err:
# 	if err.response['Error']['Code'] == 'ConditionalCheckFailedException':
# 		try:
# 			user_table.update_item(
# 				Key={
# 					'uId': author
# 				},
# 				UpdateExpression='SET uPst = :pId',
# 				ExpressionAttributeValues={
# 					':pId': [pId]
# 				},
# 			)
# 		except ClientError as err:
# 			return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'User table too busy'})}
# 	else:
# 		return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'User table too busy'})}

#################################################################################################
#TRACKING

# try:
# 	tracking_table.update_item(
# 		Key = {
# 		'eCdId': event_code + '_' + str(lhsh),
# 		'dt': created_at_day.strftime("%Y-%m-%d-%H")
# 		},
# 		UpdateExpression='ADD cnt :inc',
# 		ExpressionAttributeValues={
# 			':inc': 1
# 		}
# 		)
# except ClientError as err:
# 	print(err)

#################################################################################################
#EDGE ITEM

# try:
# 	edge_item_put = {
# 		'eCdById': event_code + '_' + author,
# 		'cTp': created_at_milliseconds,
# 		'abtId': pId,
# 		'loc': location,
# 		'vtSr': 0 
# 	}
# 	edge_table.put_item(Item=_dynamo_sanitize(edge_item_put))
# except ClientError as err:
# 	print(err)
# 	return {'statusCode': 404, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': 404, 'message': 'Edge table too busy'})}


#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$


#DYNAMO TABLE INFO:
#################################################################################################
#Post TABLE:
#ATTRIBUTES:
#key: pId (post id)

#b: cTp (used as control for event activity table, created timestamp)
#c: eCd (used as control for event activity table, event code with random int 0-11 at end for querying hottest posts everywhere)
#d: vlSr (used as control for event activity table, value of post based on voting or watch score tbd)
#e: pkVlSr (peak value score)
#f: acsCd (access code. options are DEL APV RJC RVW ADMRVW ADMAPV ADMRJC BAN)
#g: rplSr (reply score)
#h: shrSr (share score)
#i: vtSr (vote score)
#j: rpySr (replay score)
#k: bktKy (bucket key)
#k: cptn (caption)
#n: vtTyp (vote type assigned if applicable)
#n: vtCnt (vote content assigned if applicable)
#o: mTp (modified timestamp)
#p: l10Id (level 10 cell)
#q: l11Id (level 11 cell)
#r: l12Id (level 12 cell)
#s: l13Id (level 13 cell)
#t: l14Id (level 14 cell)
#u: l15Id (level 15 cell)
#v: l16Id (level 16 cell)
#v: lNm (belongs to location name if applicable)
#v: lAls (belongs to location name if applicable)
#v: ctyNm (belongs to location name if applicable)
#w: uId (user posting id)
#y: atchm (attachments: map)


#-----------------------------------------------------------------------------------------------
#(@@@POTENTIAL@@@)GSI1: 
#description: query top posts of all time near this location. walk around and find 8 encompassing level 10 cells.
#key: l10Id 
#range: vtSr
#-----------------------------------------------------------------------------------------------
#(@@@POTENTIAL@@@)GSI2: 
#description: query top posts of all time near this location. walk around and find 8 encompassing level 10 cells.
#key: lNm
#range: vtSr
#-----------------------------------------------------------------------------------------------

#################################################################################################

#EventLocal TABLE:
#ATTRIBUTES:
#key: eCdLId (EACH POST GETS 7 ENTRIES IN THIS TABLE for each level id, event code with no random int and level id)
#range: cTp (created timestamp)

#d: vlSr (value score)
#e: acsCd (access code. options are DEL APV RJC RVW ADMRVW ADMAPV ADMRJC)
#f: pId (post id)
#g: uId (poster user id)

#GSI 1 and 2 are GSI's because 10GB DynamoDB limit on LSI item collections
#-----------------------------------------------------------------------------------------------
#LSI1: 
#description: for moderators to approve posts in their area.
#key: eCdLId 
#range: acsCd
#-----------------------------------------------------------------------------------------------
#LSI2: 
#description: find hottest posts in area.
#key: eCdLId 
#range: vlSr
#-----------------------------------------------------------------------------------------------

#################################################################################################

#ActivityAll TABLE:
#ATTRIBUTES:
#key: eCd (random int 0-30)
#range: cTp (created timestamp)

#d: vlSr (value score)
#e: acsCd (access code. options are DEL APV RJC RVW ADMRVW ADMAPV ADMRJC)
#f: pId (post id)
#g: uId (poster user id)

#GSI 1 and 2 are GSI's because 10GB DynamoDB limit on LSI item collections
#-----------------------------------------------------------------------------------------------
#LSI1: 
#description: find hottest posts from everywhere.
#key: eCd 
#range: vlSr
#-----------------------------------------------------------------------------------------------
#LSI2: 
#description: for me (as a single person) to approve stuff until I find moderators
#key: eCd 
#range: acsCd (includes both RVW and ADMRVW)
#-----------------------------------------------------------------------------------------------


#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
#SIDENOTE:
#To avoid hotkeys the following implementation is done to evaluate hottest posts everywhere
#Chose a random int between 00-30 device side (could be more considering volume of posts). Lets say its 4.
#
#if there is no entry for "pstvid4"
#query the eCd "pstvid4" and retrieve the posts sorted by value from the GSI index where pk: eCd and sk: valScr.
#lets say the posts ranged from 2859 to 2798 add a new entry to the existing dictionary ["pstvid4": 2798]
#
#else:
#check the lowest value score stored in the dictionary corresponding to entry "pstvid4". Lets say its currently [pstvid4": 2798, "pstvid6": 2445]
#We then query "eCd" = "pstvid4" WHERE valScr < 2798 from the GSI index where pk: eCd and sk: valScr.
#update the lowest score on the entry with the new results. Lets say results ranged from 2798 to 2668, the dictionary entry should be ["pstvid4": 2668, "pstvid6": 2445]
#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

#################################################################################################

#ChildEvent TABLE:
#ATTRIBUTES:
#Key: pId (post id)
#Range: cTp (created timestamp)

#c: vtSr (vote score)
#d: acsCd (access code. options are DEL APV REJ RVW. default is APV for child)
#e: cId (this child's id)
#f: cTxt (child's text if applicable)
#g: mTp (modified timestamp)
#h: uId (posting user id)
#k: bktKy (bucket key if applicable)
#j: eCd (event code)


#-----------------------------------------------------------------------------------------------
#GSI1: 
#description: find top
#key: pId 
#range: vtSr
#-----------------------------------------------------------------------------------------------
#(@@@POTENTIAL GSI@@@)GSI2: 
#description: find specific child.
#key: cId 
#-----------------------------------------------------------------------------------------------


#################################################################################################

#Edge TABLE:
#key: eCdById (event code with NO random int and uId that authored event. ex: pstlikby#id or pstrplby#id)
#range: cTp (created timestamp)

#c: abtId (id, event code implicit in eCdById key)
#d: eCdFrId (for uId)
#e: lId (level 16 cell of user when action was preformed)
#f: vtSr (Vote score if applicable)

#-----------------------------------------------------------------------------------------------
#LSI1: 
#description: find if a user has done a specific action on a post
#key: eCdById 
#range: abtId 
#-----------------------------------------------------------------------------------------------
#LSI2: 
#description: if you want to evaluate relationships between users
#this is what you use.
#key: eCdById 
#range: eCdFrId 
#-----------------------------------------------------------------------------------------------
#LSI3: 
#description: if you want to evaluate user event code by vote score
#key: eCdById 
#range: vtSr 
#-----------------------------------------------------------------------------------------------


#################################################################################################

#Streams running on this table

#Notification Table:
#Key: uId (the user id to be notified)
#Range: cTp

#c: nfTxt (notification text)
#d: eCdAbtId (the id that the event corresponds to)


#################################################################################################
#EventTrack Table:
#Key: eCdId (with randint 00-30)
#Range: dt (year-month-day-hour)

#################################################################################################



#LocationAll TABLE:
#ATTRIBUTES:
#key: hsh (hashed int in range tbd from location level 10 index)
#range: lId (location id)

#b: gb (geoblocked)
#c: hrDsy (hourly density)
#e: mthDsy (month density)
#g: lNm (location name)
#g: lAls (location alias)
#g: ctyNm (appx name of city)
#h: rnd (random integer 00-30)

#-----------------------------------------------------------------------------------------------
#LSI1: 
#description: evaluate popular areas by hour
#key: hsh
#range: hrDsy
#-----------------------------------------------------------------------------------------------
#LSI2: 
#description: evaluate popular areas by month
#key: hsh
#range: mthDsy
#-----------------------------------------------------------------------------------------------
#LSI3: 
#description: location name
#key: hsh
#range: lNm
#-----------------------------------------------------------------------------------------------
#LSI4: 
#description: location alias
#key: hsh
#range: lAls
#-----------------------------------------------------------------------------------------------
#LSI5: 
#description: city name
#key: hsh
#range: ctyNm

#################################################################################################

#(TTL AND STREAMS RUNNING ON THIS TABLE)


#LocationSpecific TABLE:
#ATTRIBUTES:
#key: lId (location id)
#range: cTp (created timestamp, in milliseconds)


#b: eTp (expiring timestamp, in seconds)
#c: hr (running hour density, +1 every time post, -1 every delete from ttl, updates controls in locationall every hour)
#d: mth (running month density, +1 every time post, -1 every delete from ttl, updates controls in locationall every hour)



#################################################################################################

#User Table:
#Key: uId

#c: bnEnd (ban end timestamp)
#d: ctryCd (country code)
#e: usrnm (username)
#f: phnNb (phone number)
#g: krmaSr (karma score)
#h: mTp (modified timestamp)
#i: pswd (password)
#j: pmL16Id (permanent location level 16)
#k: pmLng (permanent longitude)
#l: pmLat (permanent latitude)
#j: pmCtyNm (permanent city name)
#j: pmCmpNm (permanent campus name)
#j: pmCmpAl (permanent campus alias)
#n: bio (biography)
#o: bDy (birthday)
#o: nme (name)
#o: pfBkKy (profile picture bucket key)
#p: acsCdRInt (options are USR, MOD, ADM with random integer 00-30)
#q: uOpts (user options: map)
#r: uRpl (user replies list)
#g: uPst (user posts list)
#a: uRpst (user reposts list)
#f: uStk (user stickers list)
#g: lFlwCt (locations following count)
#g: uFrdCt (users friend count)
#h: uFrd (users friends list)
#h: uFrdRqSt (users friend request sent)
#h: uFrdRqRc (users friend request recieved)
#g: lFlw (locations following list)
#f: uBlk (user blocking list)
#g: uTph (user trophies list)
#g: rnd (0-30)

#-----------------------------------------------------------------------------------------------
#GSI1: 
#description: retrieve user by phone number
#key: phnNb
#-----------------------------------------------------------------------------------------------
#GSI2: 
#description: retrieve user by user name
#key: rnd
#range: usrnm

#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
#$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

#POPULATION BASED RANKING FUNCTION

# def ranking_population(population, ups, downs, created, now=int(time.time()//1000)):
#   population = float(population)
#   if population > 45000.:
#       population = 450000.
#   elif population < 3.:
#       population = 3.
#   # a = (1.00001-1.0)/((log(3.)-log(45000.0)))
#   # b = 1.00001 - a * log(3.)
#   a = -1.03995448127e-05
#   b = 1.0001114250677277
#   return pow(hot(ups, downs), a * log(population) + b)

