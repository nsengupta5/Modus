from django.shortcuts import render, redirect
from django.conf import settings
from spotipy import oauth2

# Create your views here.
def spotify_auth(request):
    sp_oauth = oauth2.SpotifyOAuth(
        client_id=settings.SPOTIFY_CLIENT_ID,
        client_secret=settings.SPOTIFY_CLIENT_SECRET,
        redirect_uri=settings.SPOTIFY_REDIRECT_URI,
    )
    auth_url = sp_oauth.get_authorize_url()
    return redirect(auth_url)

def spotify_callback(request):
    code = request.GET.get('code', None)
    sp_oauth = oauth2.SpotifyOAuth(
        client_id=settings.SPOTIFY_CLIENT_ID,
        client_secret=settings.SPOTIFY_CLIENT_SECRET,
        redirect_uri=settings.SPOTIFY_REDIRECT_URI,
    )
    token_info = sp_oauth.get_access_token(code)
    access_token = token_info['access_token']
    return redirect('spotify:success')

def spotify_success(request):
    return render(request, 'spotify/success.html')
