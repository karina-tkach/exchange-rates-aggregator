<?php

declare(strict_types=1);

namespace App\Modules\Auth\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Modules\Auth\Domain\User;
use App\Modules\Auth\Http\Requests\LoginRequest;
use App\Modules\Auth\Http\Requests\RegisterRequest;
use Illuminate\Support\Facades\Auth;
use Tymon\JWTAuth\Exceptions\JWTException;
use Tymon\JWTAuth\Facades\JWTAuth;

class AuthController extends Controller
{
    public function login(LoginRequest $request)
    {
        $validatedRequest = $request->validated();

        if (!$token = Auth::guard('api')->attempt($validatedRequest)) {
            return response()->json(['message' => 'Invalid credentials'], 401);
        }

        return $this->respondWithToken($token);
    }

    public function register(RegisterRequest $request)
    {
        $validatedRequest = $request->validated();

        $user = User::query()
            ->create([
            'name' => $validatedRequest['name'],
            'email' => $validatedRequest['email'],
            'password' => bcrypt($validatedRequest['password']),
        ]);

        $token = JWTAuth::fromUser($user);
        return $this->respondWithToken($token);
    }

    public function logout()
    {
        try {
            Auth::guard('api')->logout();
        } catch (JWTException $e) {
            return response()->json(['message' => 'Failed to invalidate token'], 500);
        }

        return response()->json(['message' => 'Successfully logged out']);
    }

    public function refresh()
    {
        try {
            $token = Auth::guard('api')->refresh();
        } catch (JWTException $e) {
            return response()->json(['message' => 'Token cannot be refreshed'], 401);
        }

        return $this->respondWithToken($token);
    }

    protected function respondWithToken($token)
    {
        return response()->json([
            'access_token' => $token,
            'token_type'   => 'bearer',
            'expires_in'   => auth()->factory()->getTTL() * 60
        ]);
    }
}
