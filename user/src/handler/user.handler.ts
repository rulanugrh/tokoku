import { Request, Response } from "express";
import bcrypt from 'bcrypt';
import { prisma } from "..";
import { Payload, generateToken } from "../middleware/jwt";

export class Result {
    fname: string
    lname: string
    username: string
    avatar: string
    roleID: number
}

export const registerUser = async (req: Request, res: Response) => {
    try {
        const { fname, lname, username, email, avatar, roleID, password } = req.body
        const hashPassword = bcrypt.hash(password, 14)

        const find = await prisma.user.findUnique({
            where: {
                email: email
            }
        })

        if (find) {
            res.status(400).json({ msg: 'sorry your email have been created', code: 400 })
        }

        await prisma.user.create({
            data: {
                fname,
                lname,
                username,
                email,
                avatar,
                roleID,
                hashPassword
            }
        })

        const result: Result = new Result()
        result.avatar = avatar
        result.fname = fname
        result.lname = lname
        result.roleID = roleID
        result.username = username

        res.status(200).json({
            msg: 'success create account',
            code: 200,
            data: result,
        })

    } catch (error) {
        res.status(500).json({
            msg: 'something error',
            err: error,
        })
    }
}

export const loginUser = async (req: Request, res: Response) => {
    try {
        const { email, password } = req.body
        const find = await prisma.user.findUnique({
            where: {
                email: email
            }
        })

        if (!find) {
            res.status(404).json({ msg: 'sorry your email not found', code: 404 })
        }

        const verify = bcrypt.compare(password, find[0].password)
        if (!verify) {
            res.status(400).json({
                msg: 'sorry your password not matched',
                code: 400
            })
        }

        const payload = new Payload()
        payload.avatar = find[0].avatar
        payload.roleID = find[0].roleID
        payload.username = find[0].username

        const token = await generateToken(payload)

        res.cookie('authorization', token, {
            httpOnly: true,
            sameSite: "lax",
            maxAge: 24 * 60 * 60 * 1000
        })

        res.status(200).json({ msg: 'success login' })
    } catch (error) {
        res.status(500).json({
            msg: 'something error',
            err: error,
        })
    }
}

export const findID = async (req: Request, res: Response) => {
    try {
        const { id } = req.params
        const result = await prisma.user.findUnique({
            where: {
                id: Number(id)
            }
        })
        
        if (!result) {
            res.status(404).json({
                msg: 'user not found'
            })
        }

        const response: Result = new Result()
        response.avatar = result[0].avatar
        response.fname = result[0].fname
        response.lname = result[0].lname
        response.roleID = result[0].roleID
        response.username = result[0].username


        res.status(200).json({
            msg: 'user found',
            code: 200,
            data: response
        })
        
    } catch (error) {
        res.status(500).json({
            msg: 'something error',
            err: error,
        })
    }
}
